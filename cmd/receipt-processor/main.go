package main

import (
	"log"
	"net/http"

	"github.com/fetch-challenge/src/receipt-processor/api"
)

func main() {

	// TODO: Setup logging here

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// TODO: Return not found 404
	})
	http.Handle("/receipts/process", api.ReceiptsProcessHandler)
	http.Handle("/recipts/{id}/points", api.ReceiptPointsHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
TODO: Implement the server here
- Rate limiting
- Logging


*/
