-- Create extension pgcrypto if it does not already exist
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- Create merchants table if it does not already exist
CREATE TABLE IF NOT EXISTS restaurante.merchants (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    latitude VARCHAR(255) NOT NULL,
    longitude VARCHAR(255) NOT NULL,
    availability_radius VARCHAR(255) NOT NULL,
    open_hour TIME,
    close_hour TIME,
    rating DECIMAL(3,2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);
