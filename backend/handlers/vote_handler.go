package handlers

import (
	"cats-or-dogs/db"
	"cats-or-dogs/models"
	"encoding/json"
	"log"
	"net/http"
)

// VoteHandler handles vote-related requests
type VoteHandler struct {
	db *db.Database
}

// NewVoteHandler creates a new VoteHandler
func NewVoteHandler(database *db.Database) *VoteHandler {
	return &VoteHandler{
		db: database,
	}
}

// HandleVote handles the API endpoint for submitting votes
func (h *VoteHandler) HandleVote(w http.ResponseWriter, r *http.Request) {
	// Set response header
	w.Header().Set("Content-Type", "application/json")

	// Parse request body
	var vote models.Vote
	err := json.NewDecoder(r.Body).Decode(&vote)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate vote
	if err := vote.Validate(); err != nil {
		log.Printf("Invalid vote data: %v", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Check if email has already voted
	exists, err := h.db.CheckEmailExists(vote.Email)
	if err != nil {
		log.Printf("Error checking if email exists: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Error processing your vote")
		return
	}

	if exists {
		respondWithError(w, http.StatusConflict, "You have already voted")
		return
	}

	// Save vote to database
	err = h.db.SaveVote(vote.Name, vote.Email, vote.Choice)
	if err != nil {
		log.Printf("Error saving vote: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Error processing your vote")
		return
	}

	// Return success response
	respondWithJSON(w, http.StatusCreated, map[string]string{
		"message": "Vote submitted successfully",
	})
}

// Helper function to send error response
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

// Helper function to send JSON response
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
