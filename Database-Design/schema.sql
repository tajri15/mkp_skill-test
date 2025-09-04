-- Ekstensi untuk menghasilkan UUID secara otomatis
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Tabel untuk menyimpan data pengguna
-- Role 'customer' untuk pembeli tiket, 'admin' untuk manajemen jadwal
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    full_name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL CHECK (role IN ('customer', 'admin')),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Tabel untuk data film
CREATE TABLE movies (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    duration_minutes INT NOT NULL,
    release_date DATE,
    poster_image_url VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Tabel untuk data cabang bioskop
CREATE TABLE cinemas (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    city VARCHAR(100) NOT NULL,
    address TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Tabel untuk data studio/teater di dalam bioskop
CREATE TABLE theaters (
    id SERIAL PRIMARY KEY,
    cinema_id INT NOT NULL REFERENCES cinemas(id) ON DELETE CASCADE,
    name VARCHAR(50) NOT NULL,
    capacity INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Tabel untuk jadwal tayang film
CREATE TABLE showtimes (
    id SERIAL PRIMARY KEY,
    movie_id INT NOT NULL REFERENCES movies(id) ON DELETE CASCADE,
    theater_id INT NOT NULL REFERENCES theaters(id) ON DELETE CASCADE,
    start_time TIMESTAMP WITH TIME ZONE NOT NULL,
    end_time TIMESTAMP WITH TIME ZONE NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Tabel untuk pesanan/booking tiket
CREATE TABLE bookings (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id),
    showtime_id INT NOT NULL REFERENCES showtimes(id),
    total_price NUMERIC(12, 2) NOT NULL,
    status VARCHAR(20) NOT NULL CHECK (status IN ('pending', 'confirmed', 'cancelled', 'refunded')),
    booking_time TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Tabel untuk detail kursi yang dipesan dalam satu booking
CREATE TABLE booking_details (
    id SERIAL PRIMARY KEY,
    booking_id UUID NOT NULL REFERENCES bookings(id) ON DELETE CASCADE,
    seat_row VARCHAR(5) NOT NULL,
    seat_number INT NOT NULL,
    UNIQUE(booking_id, seat_row, seat_number)
);

-- Contoh data awal (opsional) untuk testing
-- Membuat user admin untuk login
-- Passwordnya adalah 'admin123'
INSERT INTO users (full_name, email, password_hash, role) VALUES
('Admin Bioskop', 'admin@bioskop.com', '$2a$10$fWJ.d.q.m5L5w.Xf8J.d.eF5u5Y9c3z1f6L9g4m2n7h8o.P2q.S.K', 'admin');

-- Contoh movie, cinema, theater
INSERT INTO movies (title, duration_minutes) VALUES ('All for one', 125);
INSERT INTO cinemas (name, city) VALUES ('Cinema XXI', 'Semarang');
INSERT INTO theaters (cinema_id, name, capacity) VALUES (1, 'Studio 1', 150);