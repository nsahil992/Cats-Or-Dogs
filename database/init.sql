-- Create database
CREATE DATABASE catsordogs;

-- Connect to database
\c catsordogs;

-- Create votes table
CREATE TABLE IF NOT EXISTS votes (
                                     id SERIAL PRIMARY KEY,
                                     name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    choice VARCHAR(10) NOT NULL CHECK (choice IN ('cat', 'dog')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

-- Create index on email for faster lookups
CREATE INDEX IF NOT EXISTS idx_votes_email ON votes(email);

-- Create index on choice for faster stats queries
CREATE INDEX IF NOT EXISTS idx_votes_choice ON votes(choice);