# UI Design Guidelines

This document outlines the UI design system used in the application, extracted from the current React Native (Expo) codebase. Use this guide to implement a consistent look and feel in other applications.

## 1. Design Tokens

### 1.1 Colors

The application supports both Light and Dark modes.

| Token | Light Mode Value | Dark Mode Value | Description |
| :--- | :--- | :--- | :--- |
| **Primary** | `#f59e0b` (Amber 500) | `#f59e0b` | Brand color (Buttons, Active States, Highlights) |
| **Success** | `#10b981` | `#10b981` | Success states |
| **Error** | `#ef4444` | `#ef4444` | Error states, Like Button (active) |
| **Background** | `#f8fafc` (Slate 50) | `#000000` (True Black) | Screen background |
| **Card** | `#ffffff` | `#18181b` (Zinc 900) | Surface color for cards, modals |
| **Text** | `#0f172a` (Slate 900) | `#FFFFFF` | Primary text color |
| **Text Muted** | `#64748b` (Slate 500) | `#a1a1aa` (Zinc 400) | Secondary text, captions, subtitles |
| **Border** | `#e2e8f0` (Slate 200) | `#27272a` (Zinc 800) | Borders, dividers |
| **Icon Active** | `#0f172a` | `#fbbf24` (Amber 400) | Active tab icons, primary icons |
| **Icon Inactive**| `#94a3b8` | `#475569` | Inactive tab icons |

**Special Semantic Colors:**
*   **Rating Star:** `#fbbf24` (Amber 400)
*   **Slot Icon Bg:** `#fffbeb` (Light) / `rgba(245, 158, 11, 0.1)` (Dark)
*   **Tab Bar Bg:** `rgba(255, 255, 255, 0.95)` (Light) / `rgba(15, 23, 42, 0.95)` (Dark)

### 1.2 Spacing & Layout

The system uses a 4-point grid system.

*   **xs:** 4px
*   **sm:** 8px
*   **md:** 12px
*   **lg:** 16px
*   **xl:** 20px
*   **xxl:** 24px

### 1.3 Border Radius

Rounded corners are heavily used to create a friendly, modern feel.

*   **sm:** 8px (Tags, Small Buttons)
*   **md:** 12px (Inner Cards, Badges)
*   **lg:** 16px (Card Containers, Avatars)
*   **xl:** 24px (Large Containers, Search Bars)
*   **full:** 9999px (Circle Buttons, Capsules)

### 1.4 Shadows (Light Mode Only)

Shadows are generally removed or replaced with borders in Dark Mode.

*   **Card Shadow:**
    *   Offset: `{ width: 0, height: 8 }`
    *   Radius: `10`
    *   Opacity: `0.1`
    *   Color: `#000000`
    *   Elevation: `4` (Android)
*   **Button Shadow:**
    *   Offset: `{ width: 0, height: 4 }`
    *   Radius: `10`
    *   Opacity: `0.1`
    *   Color: `#0f172a`

---

## 2. Typography

The application relies on the system font stack (`System`, `San Francisco`, `Roboto`).
*   **Family:** System Default (Fallback to sans-serif)

**Hierarchy:**

| Style | Size | Weight | Line Height | Case | Usage |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **Heading 1** | 18px | Bold (700) | Auto | Title Case | Screen Titles, Card Titles |
| **Body** | 14px | Medium (500) | Auto | Sentence | Input text, Main content |
| **Subtitle** | 12px | Medium (500) | Auto | Sentence | Location, Metadata |
| **Label** | 10-12px | Bold (700) | Auto | Uppercase | Tags, Slot Labels, Tab Labels |
| **Caption** | 10px | Medium (500) | Auto | Sentence | Reviews count |

**Letter Spacing:**
*   Labels/Tags often use `letterSpacing: 1` for uppercase text.

---

## 3. Iconography

*   **Library:** `lucide-react-native`
*   **Stroke Width:**
    *   Standard: `2`
    *   Active/Selected: `2.5`
*   **Sizing:**
    *   Small (Metadata): `12px - 14px`
    *   Medium (Buttons, Inputs): `16px - 18px`
    *   Large (Tab Bar): `22px`
    *   X-Large (Empty States): `48px`

---

## 4. Components

### 4.1 Shop Card
A complex card displaying a shop image, details, and booking action.

*   **Container:** `borderRadius: 24`, `borderWidth: 1`.
*   **Image:** Aspect ratio approx 16:9 (height 160px), `borderRadius: 16`.
*   **Overlays:**
    *   **Favorite Button:** Top-right, Circular (36x36), translucent black bg.
    *   **Rating Badge:** Top-left, translucent black bg, rounded.
*   **Footer:** Dashed top border (`borderStyle: 'dashed'`), separates content from "Book Now" button.

### 4.2 Buttons
*   **Primary Button:**
    *   Bg: Primary Color (`#f59e0b` or `#0f172a` depending on context)
    *   Radius: `12px` (md)
    *   Text: Bold, 12px
    *   Padding: Vertical 10px, Horizontal 16px
*   **Icon Button (Circle):**
    *   Size: 36x36
    *   Radius: Full
    *   Used for: Favorites, Back buttons, Slot icons.

### 4.3 Inputs & Search
*   **Search Bar:**
    *   Height: ~40-44px
    *   Radius: `20px` (xl)
    *   Background: Surface color (`#ffffff` / `#000000`)
    *   Border: 1px solid Border color
    *   Icon: Left-aligned search icon.

### 4.4 Filter Chips
*   **Container:** Capsule shape (`borderRadius: 12-16`).
*   **State - Default:** Border 1px, Surface Background.
*   **State - Active:** Primary Background, White Text, No Border.

### 4.5 Navigation (Tab Bar)
*   **Position:** Bottom floating (visually), with blur/opacity effect.
*   **Background:** 95% opacity.
*   **Interaction:**
    *   Active tab has a pill-shaped indicator background (`opacity: 1`, `scale: 1`).
    *   Icons scale slightly.
    *   Labels are small (10px).

---

## 5. Interactions & Animations

*   **Touch Feedback:**
    *   **ScalePress:** Most interactive elements (Cards, Buttons, Chips) use a scaling animation on press.
    *   **Scale Down:** to `0.96`.
    *   **Physics:** Spring animation (`speed: 20`, `bounciness: 10`).
*   **Theme Transition:**
    *   Smooth color transitions for background and icon toggles.
    *   Toggle switch uses a spring translation animation (`translateX`).
*   **Skeleton Loading:**
    *   Used for cards while data is fetching.
    *   Pulsing opacity animation.
