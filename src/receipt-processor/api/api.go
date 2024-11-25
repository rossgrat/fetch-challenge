package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/rossgrat/fetch-challenge/src/logger"
	"github.com/rossgrat/fetch-challenge/src/receipt-processor/db"
)

type ErrorResponse struct {
	Description string
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}
type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}
type ReceiptID struct {
	ID string `json:"id"`
}

// /receipts/process
// In this endpoints, we receive a Receipt object, score it for points,
// then save those points as a valid receipt UUID, returing the UUID
//
// Note: We do not save receipts, or any kind of signature in the
// database, if we wanted to prevent duplicate receipts from being sent,
// we would do that here
func ReceiptsProcessHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		WriteResponse(w, r, http.StatusNotFound,
			ErrorResponse{Description: "Incorrect method"})
		return
	}

	var receipt Receipt
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&receipt); err != nil {
		logger.LogInfo(r, err.Error())
		WriteResponse(w, r, http.StatusBadRequest,
			ErrorResponse{Description: "The receipt is invalid"})
		return
	}

	if err := ValidateReceipt(receipt); err != nil {
		logger.LogInfo(r, err.Error())
		WriteResponse(w, r, http.StatusBadRequest,
			ErrorResponse{Description: "The receipt is invalid"})
		return
	}

	points := CalculateReceiptPoints(receipt)

	dbReceipt := db.Receipt{
		Points: points,
	}
	receiptID, err := db.CreateReceipt(dbReceipt)
	if err != nil {
		logger.LogInfo(r, err.Error())
		WriteResponse(w, r, http.StatusBadRequest,
			ErrorResponse{Description: "The receipt is invalid"})
		return
	}

	response := ReceiptID{
		ID: receiptID,
	}
	WriteResponse(w, r, http.StatusOK, response)
}

type ReceiptPoints struct {
	Points int `json:"points"`
}

// receipts/:id/points
func ReceiptPointsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		WriteResponse(w, r, http.StatusNotFound,
			ErrorResponse{Description: "Incorrect method"})
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	id := pathParts[2] // Router ensures path has at least 3 splits

	dbReceipt, err := db.GetReceipt(id)
	if err != nil {
		WriteResponse(w, r, http.StatusNotFound,
			ErrorResponse{Description: "No receipt found for that id"})
		return
	}

	response := ReceiptPoints{
		Points: dbReceipt.Points,
	}
	WriteResponse(w, r, http.StatusOK, response)
}
