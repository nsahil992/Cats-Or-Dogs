package db

import (
	"cats-or-dogs/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Database represents a connection to the PostgreSQL database
type Database struct {
	*sql.DB
}

// Connect establishes a connection to the PostgreSQL database
func Connect(cfg config.DBConfig) (*Database, error) {
	// Build connection string
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	// Open database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %w", err)
	}

	// Verify connection is working
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	log.Println("Connected to database successfully")

	// Initialize database schema
	if err := initSchema(db); err != nil {
		return nil, fmt.Errorf("error initializing database schema: %w", err)
	}

	return &Database{db}, nil
}

// initSchema creates necessary tables if they don't exist
func initSchema(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS votes (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) NOT NULL,
		choice VARCHAR(10) NOT NULL CHECK (choice IN ('cat', 'dog')),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	log.Println("Database schema initialized")
	return nil
}

// SaveVote inserts a new vote into the database
func (db *Database) SaveVote(name, email, choice string) error {
	query := `
	INSERT INTO votes (name, email, choice)
	VALUES ($1, $2, $3)
	`

	_, err := db.Exec(query, name, email, choice)
	if err != nil {
		return fmt.Errorf("error saving vote to database: %w", err)
	}

	return nil
}

// GetStats retrieves the count of votes for cats and dogs
func (db *Database) GetStats() (map[string]int, error) {
	stats := map[string]int{
		"cats": 0,
		"dogs": 0,
	}

	// Count cats
	catQuery := `SELECT COUNT(*) FROM votes WHERE choice = 'cat'`
	var catCount int
	err := db.QueryRow(catQuery).Scan(&catCount)
	if err != nil {
		return nil, fmt.Errorf("error getting cat vote count: %w", err)
	}
	stats["cats"] = catCount

	// Count dogs
	dogQuery := `SELECT COUNT(*) FROM votes WHERE choice = 'dog'`
	var dogCount int
	err = db.QueryRow(dogQuery).Scan(&dogCount)
	if err != nil {
		return nil, fmt.Errorf("error getting dog vote count: %w", err)
	}
	stats["dogs"] = dogCount

	return stats, nil
}

// CheckEmailExists checks if an email has already voted
func (db *Database) CheckEmailExists(email string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM votes WHERE email = $1)`

	var exists bool
	err := db.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking if email exists: %w", err)
	}

	return exists, nil
}
