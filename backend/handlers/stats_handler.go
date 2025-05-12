package handlers

import (
	"cats-or-dogs/db"
	"log"
	"net/http"
)

// StatsHandler handles statistics-related requests
type StatsHandler struct {
	db *db.Database
}

// NewStatsHandler creates a new StatsHandler
func NewStatsHandler(database *db.Database) *StatsHandler {
	return &StatsHandler{
		db: database,
	}
}

// HandleStats handles the API endpoint for retrieving vote statistics
func (h *StatsHandler) HandleStats(w http.ResponseWriter, r *http.Request) {
	// Set response header
	w.Header().Set("Content-Type", "application/json")

	// Get stats from database
	stats, err := h.db.GetStats()
	if err != nil {
		log.Printf("Error retrieving stats: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Error retrieving statistics")
		return
	}

	// Return stats
	respondWithJSON(w, http.StatusOK, stats)
}
