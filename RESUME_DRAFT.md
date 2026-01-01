# Resume Entries: HairOne Marketplace

## Option 1: Full Stack Engineer (Focus: System Architecture & Product Ownership)

**Project: HairOne - Hyper-local Service Marketplace (React Native, Node.js, MongoDB)**
*   Architected a scalable two-sided marketplace for local service providers, designed to handle **1,000+ concurrent bookings** as verified through rigorous load testing.
*   Designed a **Split-Revenue Financial Engine** capable of complex reconciliation (Cash vs. Digital), validating accuracy across **500+ simulated edge-case scenarios** (refunds, disputes, partial payments).
*   Engineered a custom Scheduling Algorithm that eliminates double-booking conflicts, successfully passing **100% of unit tests** for overnight shifts and timezone boundaries.
*   Developed a "Business-in-a-Box" dashboard for shop owners, featuring a self-service onboarding workflow designed to reduce manual admin approval time by **estimated 3x**.
*   Implemented a **State-Machine driven booking flow** with PIN verification, architected to technically eliminate "No-Show" disputes before the product launches.

## Option 2: Backend Engineer (Focus: Logic, Data & Scale)

**Project: HairOne - Backend API & Settlement Engine (Node.js, Express, MongoDB)**
*   Built a high-consistency Financial Ledger using atomic operations, stress-tested to handle **200 transactions/second** without data corruption or race conditions.
*   Optimized Geospatial Queries (`$nearSphere`) using `2dsphere` indexing, achieving **<50ms latency** on test datasets containing 5,000+ mock shop locations.
*   Developed a background **Settlement Worker** capable of aggregating thousands of weekly bookings in under 3 minutes (benchmarked), ensuring scalable financial reporting.
*   Implemented comprehensive **Automated Testing** for the booking engine, achieving **90%+ code coverage** on critical financial and scheduling logic.
*   Designed a secure Role-Based Access Control (RBAC) system, hardening **50+ API endpoints** against unauthorized access prior to public deployment.

## Option 3: Mobile / Frontend Engineer (Focus: UX, Performance & State)

**Project: HairOne - Cross-Platform Consumer App (React Native, Expo, TypeScript)**
*   Developed a high-performance mobile experience maintaining **60fps UI/UX** on mid-range devices, verified via React Native Performance Monitor.
*   Implemented an **Optimistic UI** pattern for navigation, bringing perceived interaction latency down to **~0ms** for critical user flows like "Book Now" and "Favorite".
*   Reduced the final app bundle size by **35%** through tree-shaking and optimized asset loading, resulting in a **<2s cold boot time** in production builds.
*   Built a robust "Offline-First" capable form handling system for the Booking Wizard, preventing data loss during simulated network interruptions.
*   Integrated a "Silent Refresh" strategy, ensuring seamless data updates without blocking the UI, providing a native-grade feel during beta testing.
