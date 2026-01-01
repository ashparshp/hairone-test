-- Users
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    phone VARCHAR(20) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    name VARCHAR(100),
    avatar VARCHAR(255),
    email VARCHAR(100),
    gender VARCHAR(20),
    role VARCHAR(20) DEFAULT 'user', -- 'user', 'admin', 'owner'
    is_premium BOOLEAN DEFAULT FALSE,
    business_name VARCHAR(100),
    application_status VARCHAR(20) DEFAULT 'none',
    suspension_reason TEXT,
    token_version INTEGER DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Shops
CREATE TABLE shops (
    id SERIAL PRIMARY KEY,
    owner_id INTEGER NOT NULL REFERENCES users(id),
    name VARCHAR(100) NOT NULL,
    address TEXT NOT NULL,
    coordinates JSONB, -- { lat: Number, lng: Number }
    image VARCHAR(255),
    gallery JSONB DEFAULT '[]', -- Array of strings
    type VARCHAR(20) DEFAULT 'unisex',
    rating NUMERIC(3, 2) DEFAULT 0,
    review_count INTEGER DEFAULT 0,
    buffer_time INTEGER DEFAULT 0,
    min_booking_notice INTEGER DEFAULT 60,
    max_booking_notice INTEGER DEFAULT 30,
    auto_approve_bookings BOOLEAN DEFAULT TRUE,
    block_custom_bookings BOOLEAN DEFAULT FALSE,
    is_disabled BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Services
CREATE TABLE services (
    id SERIAL PRIMARY KEY,
    shop_id INTEGER NOT NULL REFERENCES shops(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    duration INTEGER NOT NULL, -- minutes
    is_available BOOLEAN DEFAULT TRUE
);

-- Combos
CREATE TABLE combos (
    id SERIAL PRIMARY KEY,
    shop_id INTEGER NOT NULL REFERENCES shops(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    original_price NUMERIC(10, 2) NOT NULL,
    duration INTEGER NOT NULL,
    items JSONB DEFAULT '[]', -- Array of Service IDs or Names (simplifying as JSONB for now to match Mongo structure)
    is_available BOOLEAN DEFAULT TRUE
);

-- Barbers
CREATE TABLE barbers (
    id SERIAL PRIMARY KEY,
    shop_id INTEGER NOT NULL REFERENCES shops(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    avatar VARCHAR(255),
    start_hour VARCHAR(10) DEFAULT '10:00',
    end_hour VARCHAR(10) DEFAULT '20:00',
    breaks JSONB DEFAULT '[]', -- [{startTime, endTime, title}]
    weekly_schedule JSONB DEFAULT '[]', -- Complex structure, kept as JSONB for flexibility
    special_hours JSONB DEFAULT '[]',
    is_available BOOLEAN DEFAULT TRUE
);

-- Settlements
CREATE TABLE settlements (
    id SERIAL PRIMARY KEY,
    shop_id INTEGER NOT NULL REFERENCES shops(id),
    admin_id INTEGER, -- Optional manual trigger
    type VARCHAR(20) NOT NULL, -- 'PAYOUT', 'COLLECTION'
    amount NUMERIC(10, 2) NOT NULL,
    status VARCHAR(50) DEFAULT 'PENDING',
    bookings JSONB DEFAULT '[]', -- Array of Booking IDs included
    date_range_start TIMESTAMP WITH TIME ZONE,
    date_range_end TIMESTAMP WITH TIME ZONE,
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Bookings
CREATE TABLE bookings (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id), -- Nullable for blocked slots
    shop_id INTEGER NOT NULL REFERENCES shops(id),
    barber_id INTEGER NOT NULL REFERENCES barbers(id),
    service_names JSONB DEFAULT '[]', -- Array of strings
    total_price NUMERIC(10, 2),
    total_duration INTEGER,

    date VARCHAR(20) NOT NULL, -- "YYYY-MM-DD"
    start_time VARCHAR(10) NOT NULL, -- "14:30"
    end_time VARCHAR(10) NOT NULL,

    status VARCHAR(20) DEFAULT 'upcoming', -- 'upcoming', 'completed', 'cancelled', 'pending'
    type VARCHAR(20) DEFAULT 'online',
    payment_method VARCHAR(20) DEFAULT 'PAY_AT_VENUE',
    booking_key VARCHAR(50),
    is_rated BOOLEAN DEFAULT FALSE,
    notes TEXT,

    original_price NUMERIC(10, 2),
    discount_amount NUMERIC(10, 2),
    final_price NUMERIC(10, 2),

    admin_commission NUMERIC(10, 2),
    admin_net_revenue NUMERIC(10, 2),
    barber_net_revenue NUMERIC(10, 2),

    amount_collected_by VARCHAR(20) DEFAULT 'BARBER',
    settlement_status VARCHAR(20) DEFAULT 'PENDING',
    settlement_id INTEGER REFERENCES settlements(id),

    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Reviews
CREATE TABLE reviews (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    shop_id INTEGER NOT NULL REFERENCES shops(id),
    booking_id INTEGER NOT NULL REFERENCES bookings(id),
    rating INTEGER NOT NULL,
    comment TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
