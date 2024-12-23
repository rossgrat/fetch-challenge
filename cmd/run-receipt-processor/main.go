package main

import (
	"log"
	"net/http"

	"github.com/rossgrat/fetch-challenge/src/logger"
	"github.com/rossgrat/fetch-challenge/src/mw"
	"github.com/rossgrat/fetch-challenge/src/receipt-processor/api"
)

func main() {
	http.HandleFunc("/", mw.LogRequest(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))

	http.HandleFunc("/receipts/process", mw.LogRequest(api.ReceiptsProcessHandler))
	http.HandleFunc("/receipts/{id}/points", mw.LogRequest(api.ReceiptPointsHandler))

	logger.LogInfo(nil, "Starting up!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
