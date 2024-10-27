package main

import (
	"blast_developer_challenges/handlers"
	"log"
	"net/http"
)


func main() {

	http.HandleFunc("/api/match", handlers.GetMatchStarts)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}