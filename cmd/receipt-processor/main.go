package main

import (
	"log"
	"net/http"

	"github.com/rossgrat/fetch-challenge/src/receipt-processor/api"
)

func main() {

	// TODO: Setup logging here

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	// TODO: Wrap handlers in logging middleware
	http.HandleFunc("/receipts/process", api.ReceiptsProcessHandler)
	http.HandleFunc("/receipts/{id}/points", api.ReceiptPointsHandler)

	// TODO: Configure
	log.Println("Starting up!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
TODO: Implement the server here
- Rate limiting
- Logging


*/
