package handlers

import (
	"blast_developer_challenges/parser"
	"encoding/json"
	"net/http"
)


func GetMatchStarts(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Ctronol-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Ctronol-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Ctronol-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Ctronol-Allow-Headers", "Content-Type")

	// Parse the match log and return the data as JSON
	match, err := parser.ParseMatchLog("data/NAVIvsVitaNuke.txt")
	if err != nil {
		http.Error(w, "Error parsing match data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(match)
}