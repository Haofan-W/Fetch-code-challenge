package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var (
	receipts = make(map[string]ReceiptWithPoint) //map to store receipts by ID
)

func processReciptHandler(w http.ResponseWriter, r *http.Request) {
	var receipt Receipt

	// parse JSON into receipt struct
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		log.Printf("[processReciptHandler: error parsing JSON file into receipt.\"%s\"]\n", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// calculate points for the receipt
	points, err := calculateReceiptPoints(receipt)

	if err != nil {
		log.Printf("[processReciptHandler: error calculating receipt points. \"%s\"]\n", err)
	}

	receiptId := uuid.New().String()

	receipts[receiptId] = ReceiptWithPoint{
		Receipt: receipt,
		Points:  points,
	}

	response := ReceiptOutput{
		ID:     receiptId,
		Points: points,
	}

	data, err := json.Marshal(response)
	if err != nil {
		log.Printf("[processReciptHandler: error convert response to JSON. \"%s\"]\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func getPointsHandler(w http.ResponseWriter, r *http.Request) {
	// Extract Id from URL parameter

	fmt.Println("Received request:", r.URL.Path)
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Println("the id is", id)

	receiptWithPoint, exists := receipts[id]
	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	points := receiptWithPoint.Points

	response := PointsResponse{
		Points: points,
	}

	data, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/receipts/{id}/points", getPointsHandler)
	router.HandleFunc("/receipts/process", processReciptHandler)

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", router)
}
