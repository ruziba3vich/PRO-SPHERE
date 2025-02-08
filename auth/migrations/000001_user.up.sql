DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'gender') THEN
        CREATE TYPE gender AS ENUM('male', 'female', 'none');
    END IF;
END $$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'role') THEN
        CREATE TYPE role AS ENUM('super_admin', 'admin', 'moderator', 'user');
    END IF;
END $$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status') THEN
        CREATE TYPE status AS ENUM ('active', 'inactive');
    END IF;
END $$;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    pro_id INT,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    email VARCHAR DEFAULT NULL,
    date_of_birth DATE DEFAULT NULL,
    gender gender NOT NULL,
    avatar VARCHAR DEFAULT NULL,
    phone VARCHAR NOT NULL,
    role role DEFAULT 'user',
    status status DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Add index for pro_id
CREATE UNIQUE INDEX idx_unique_pro_id ON users (pro_id);
-- Drop the index on pro_id column
