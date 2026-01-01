# Resume Entries: HairOne Marketplace

## Option 1: Full Stack Engineer (Focus: System Architecture & Product Ownership)

**Project: HairOne - Hyper-local Service Marketplace (React Native, Node.js, MongoDB)**
*   Architected and built a two-sided marketplace connecting customers with local service providers, featuring real-time availability and geospatial discovery.
*   Designed a complex **Split-Revenue Financial Engine** capable of reconciling mixed payment methods (Cash vs. Digital) and automating commission settlements via nightly Cron jobs.
*   Implemented a robust **Scheduling Algorithm** handling edge cases like overnight shifts (spillovers), time-zone handling (IST), and buffer periods to prevent double-bookings.
*   Developed a "Business-in-a-Box" dashboard for shop owners including CRM features, financial reporting (Aggregation Pipelines), and Portfolio management.
*   Engineered a **State-Machine driven booking flow** with PIN-based verification to eliminate "No-Show" disputes and ensure service delivery confirmation.

## Option 2: Backend Engineer (Focus: Logic, Data & Scale)

**Project: HairOne - Backend API & Settlement Engine (Node.js, Express, MongoDB)**
*   Built a highly consistent **Financial Ledger** that snapshots commission rates and net revenues at the moment of transaction to ensure historical data integrity during rate changes.
*   Developed a background **Settlement Worker (Cron)** that aggregates daily bookings, calculates net Debit/Credit positions for hundreds of shops, and generates automated payout reports.
*   Optimized **Geospatial Queries** (`$nearSphere`) to enable efficient radius-based search for thousands of potential shop locations with millisecond latency.
*   Implemented strict **Concurrency Controls** using atomic database operations to prevent race conditions during simultaneous booking attempts for the same time slot.
*   Designed a secure RBAC system distinguishing between Customers, Shop Owners, and Super Admins, protecting sensitive financial endpoints.

## Option 3: Mobile / Frontend Engineer (Focus: UX, Performance & State)

**Project: HairOne - Cross-Platform Consumer App (React Native, Expo, TypeScript)**
*   Developed a high-performance mobile experience using **Expo Router** and **Reanimated**, achieving 60fps animations and seamless page transitions.
*   Implemented an **Optimistic UI** pattern for the booking flow, allowing users to navigate and interact with shop details instantly while heavy data fetches occurred in the background.
*   Built a custom **Design System** with dark/light mode support, creating reusable components for complex inputs like interactive Time Slot Pickers and Image Galleries.
*   Integrated **Map Features** for location-based discovery, utilizing device geolocation to sort and filter services dynamically based on user proximity.
*   Managed complex global state for the "Booking Wizard," handling multi-step form validation, payment method selection, and real-time error handling.
