package models

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

// Vote represents a user's vote
type Vote struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Choice    string    `json:"choice"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Stats represents the total votes for cats and dogs
type Stats struct {
	Cats int `json:"cats"`
	Dogs int `json:"dogs"`
}

// Validate checks if the vote data is valid
func (v *Vote) Validate() error {
	// Validate name
	if strings.TrimSpace(v.Name) == "" {
		return errors.New("name cannot be empty")
	}

	if len(v.Name) > 100 {
		return errors.New("name is too long (maximum 100 characters)")
	}

	// Validate email
	if strings.TrimSpace(v.Email) == "" {
		return errors.New("email cannot be empty")
	}

	if !isValidEmail(v.Email) {
		return errors.New("invalid email format")
	}

	if len(v.Email) > 100 {
		return errors.New("email is too long (maximum 100 characters)")
	}

	// Validate choice
	v.Choice = strings.ToLower(strings.TrimSpace(v.Choice))
	if v.Choice != "cat" && v.Choice != "dog" {
		return errors.New("choice must be either 'cat' or 'dog'")
	}

	return nil
}

// Helper function to validate email format
func isValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(pattern, email)
	return match
}
