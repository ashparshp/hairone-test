# Interview Preparation: HairOne Marketplace

This document contains 20 detailed Q&A pairs designed to help you explain the technical depth of the HairOne project during an interview.

---

## I. High-Level & Architecture

### 1. Tell me about the "HairOne" project.
**Answer:**
"HairOne is a hyper-local marketplace I built to connect customers with local barbers and salons. Unlike simple directory apps, this handles the entire lifecycle: from finding a shop near you using geospatial search, to booking a specific 15-minute slot, to handling split payments.
Technically, it's a **React Native** app on the frontend and a **Node.js/Express** monolith on the backend with **MongoDB**.
The most complex part was the **Financial Engine**: because customers can pay via Cash (to the shop) or Online (to the platform), I had to build a reconciliation system that balances the books nightly to ensure the Platform gets its commission and the Shop gets its earnings."

### 2. Why did you choose a Monolithic architecture over Microservices?
**Answer:**
"For this stage of the product, a Monolith offered the best balance of velocity and simplicity.
1.  **Shared Types:** I could easily share logic and types between the API and the background workers.
2.  **Atomic Consistency:** Managing the financial transactions in a single database made it easier to ensure ACID-like properties without the complexity of distributed transactions.
However, I designed it with **modularity** in mind. The 'Settlement Engine' runs as a separate Cron job that could easily be peeled off into its own microservice if the load requires it in the future."

### 3. What was the most challenging technical problem you solved?
**Answer:**
"The **'Overnight Spillover'** scheduling problem.
Many shops work late (e.g., until 2 AM), which technically crosses into the next day. A naive query for 'Slots on Tuesday' would miss a shift that started Monday night and ended Tuesday morning.
**Solution:** I implemented a logic that checks two schedules for any given slot request: Today's schedule, AND Yesterday's schedule (checking if it flows over `end > 1440 minutes`). This ensured we never double-booked a late-night slot."

### 4. How did you ensure the system is scalable?
**Answer:**
"I focused on database efficiency.
1.  **Indexing:** I used Geospatial indexes (`2dsphere`) for the location search and Compound indexes for the booking lookups (`shopId` + `date` + `status`) to keep read latency low.
2.  **Aggregation:** Instead of processing financial reports in Node.js (which would block the event loop), I pushed the heavy lifting to MongoDB using **Aggregation Pipelines** to sum up thousands of records instantly."

---

## II. Backend & Logic (The "Meat")

### 5. How does your "Availability Engine" work?
**Answer:**
"It's a read-heavy operation. When a user asks for slots:
1.  I fetch the Shop's 'Weekly Schedule' and 'Special Hours' (holidays).
2.  I generate every possible 15-minute chunk for the day.
3.  I filter out slots that clash with existing bookings.
4.  **Critical Detail:** I specifically check for `bufferTime`. If a haircut is 30 mins, I not only need a 30-min gap, I need `30 mins + 5 mins cleanup`. My algorithm calculates this collision detection dynamically."

### 6. Explain the "Split-Revenue" model you mentioned.
**Answer:**
"It's based on the flow of money.
*   **Case A (Online):** Platform collects $100. Platform owes Shop $90.
*   **Case B (Cash):** Shop collects $100. Shop owes Platform $10.
I store this state on *every single booking* using an `amountCollectedBy` flag ('ADMIN' vs 'BARBER'). This immutability is key—we never calculate 'who owes whom' on the fly; we simply aggregate the stored flags."

### 7. What happens if you change the Commission Rate? Does it break old records?
**Answer:**
"No. That's a classic trap I avoided.
The booking controller takes a **Financial Snapshot** at the moment of creation. We store `adminCommission`, `originalPrice`, and `netRevenue` as fixed numbers in the Booking document.
If I change the global commission from 10% to 15% tomorrow, it only affects *future* bookings. Historical financial reports remain 100% accurate to the penny."

### 8. How do you handle the nightly settlements?
**Answer:**
"I use `node-cron` to trigger a job at 00:00.
1.  **Cutoff:** It only settles bookings that were 'Completed' before the current week started.
2.  **Aggregation:** It runs a MongoDB aggregation to group by `shopId`.
3.  **Netting:** It sums (Credits - Debits).
4.  **Result:** If the result is positive, we generate a `Payout` record. If negative, we generate a `Collection Invoice`."

### 9. How do you prevent "No-Shows"?
**Answer:**
"I implemented a **State Machine** with PIN verification.
A booking stays in `Upcoming` state. When the customer sits in the chair, the Barber asks for their **4-digit PIN** (generated at booking time).
Entering this PIN moves the state to `Checked-In`.
This solves two things:
1.  It proves the customer actually arrived (No-Show prevention).
2.  It prevents the Shop from falsely claiming a 'No-Show' to avoid paying commission."

### 10. How do you handle Time Zones?
**Answer:**
"The server is stateless, but the business is local to India.
I standardized everything to handle **IST**. Even if the server runs in UTC (like on AWS), my date utilities explicitly convert to IST before checking 'Is the shop open?'.
This prevents a bug where a booking at 1 AM UTC might accidentally be allowed for a shop that is actually closed at 6:30 AM IST."

### 11. How do you prevent Race Conditions (Two people booking the same slot)?
**Answer:**
"MongoDB doesn't have row-level locking like SQL, but it has **Atomic Operations**.
However, for complex slot checking, I rely on the single-threaded nature of Node.js for the check-and-create block.
For higher scale, I would use a `Redlock` (Redis Lock) on the `shopId + time` key during the creation process to ensure strict serialization of requests for the exact same resource."

### 12. How do you handle Image Uploads?
**Answer:**
"I use `multer-s3` to stream uploads directly to **DigitalOcean Spaces** (S3 compatible).
We never store files on the app server (which is ephemeral). We only store the public URL string in MongoDB. This keeps the backend stateless and easy to scale horizontally."

---

## III. Frontend (React Native)

### 13. How did you handle the complex Navigation?
**Answer:**
"I used **Expo Router** (file-based routing).
The trickiest part was the 'Booking Wizard'. It's a multi-step flow (Select Service -> Select Date -> Select Slot -> Confirm).
I managed this state globally but reset it on unmount to ensure that if a user abandons a booking and comes back, they start fresh."

### 14. What is "Optimistic UI" and where did you use it?
**Answer:**
"Optimistic UI is updating the interface *before* the server responds.
I used it in the **Like/Favorite** button. When you heart a shop, it turns red instantly.
I also used it for **Navigation**: When clicking a Shop Card, I pass the image and name as route params to the Details Page. This allows the 'Skeleton' of the page to render instantly with data, while the full details (reviews, services) fetch in the background."

### 15. How did you implement the "Silent Refresh"?
**Answer:**
"When a user returns to the 'My Bookings' screen, we don't want to show a loading spinner every time.
I implemented a `useFocusEffect` that triggers a data fetch.
However, I keep the *old data* visible while the new data fetches. The user only sees a change if the data is different. This makes the app feel 'native' and incredibly fast."

### 16. How did you handle Location Services?
**Answer:**
"I used `expo-location`.
1.  **Permission:** We ask for permission gracefully.
2.  **Fallback:** If denied, we default to a central city location.
3.  **Sort:** We pass the user's lat/long to the backend `getAllShops` endpoint, which uses the Haversine formula to return shops sorted by distance."

---

## IV. General & Soft Skills

### 17. How did you decide on the "Cash Limit" feature?
**Answer:**
"I noticed a potential abuse vector: Users could book 10 slots with 'Cash' payment and never show up, blocking the barber's calendar.
I introduced a **Business Logic Rule**: Users can only have 5 active/cancelled 'Cash' bookings per month. After that, they *must* pay online. This balances user flexibility with shop protection."

### 18. How do you handle Errors?
**Answer:**
"On the backend, I have a central **Error Middleware**. No matter where the error happens, it bubbles up to this middleware which formats a standard JSON response (`{ message: '...', code: 400 }`).
On the frontend, I use `Axios Interceptors`. If we get a 401 (Unauthorized), we silently try to refresh the token or redirect to Login, providing a seamless loop."

### 19. If you had 2 more weeks, what would you add?
**Answer:**
"I would add **WebSockets (Socket.io)** for real-time updates.
Right now, if a shop accepts a booking, the user has to refresh to see the status change. With WebSockets, I could push that notification instantly.
I would also add a **Waitlist** feature for fully booked days."

### 20. Why did you choose React Native over Flutter?
**Answer:**
"I come from a JavaScript/Web background. React Native allowed me to leverage my existing knowledge of React, Hooks, and State Management.
Plus, with **Expo**, the development velocity is unmatched—I can test on my physical device instantly without compiling native code every time. The ecosystem (libraries like `date-fns`, `axios`) is also much richer."
