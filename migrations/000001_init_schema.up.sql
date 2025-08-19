CREATE TABLE users (
    id serial PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(20) NOT NULL,
    driver_license VARCHAR(50) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE cars (
    id serial PRIMARY KEY,
    model VARCHAR(100) NOT NULL,
    year INTEGER NOT NULL,
    color VARCHAR(50) NOT NULL,
    mileage INTEGER NOT NULL CHECK (mileage >= 0),
    price_per_day FLOAT NOT NULL CHECK (price_per_day > 0),
    is_available BOOLEAN DEFAULT TRUE,
    location VARCHAR(255) NOT NULL,
    owner_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE rentals (
    id serial PRIMARY KEY,
    car_id int NOT NULL REFERENCES cars(id) ON DELETE CASCADE,
    user_id int NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    total_price FLOAT NOT NULL CHECK (total_price > 0),
    status VARCHAR(50) NOT NULL CHECK (status IN ('active', 'completed')),
    CONSTRAINT rentals_check CHECK (end_date > start_date)
);

CREATE TABLE reviews (
    id serial PRIMARY KEY,
    rental_id int NOT NULL REFERENCES rentals(id) ON DELETE CASCADE,
    rating INTEGER NOT NULL CHECK (rating >= 1 AND rating <= 5),
    comment TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);