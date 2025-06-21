-- This file should NOT create the DB if you're using POSTGRES_DB
-- Connect to it instead
\connect catsordogs;

CREATE TABLE IF NOT EXISTS votes (
                                     id SERIAL PRIMARY KEY,
                                     name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    choice VARCHAR(10) NOT NULL CHECK (choice IN ('cat', 'dog')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

CREATE INDEX IF NOT EXISTS idx_votes_email ON votes(email);
CREATE INDEX IF NOT EXISTS idx_votes_choice ON votes(choice);
