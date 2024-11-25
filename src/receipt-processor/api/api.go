package api

import (
	"net/http"

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
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var receipt Receipt
	if err := ReadJSONFromRequest(r, &receipt); err != nil {
		WriteResponse(w, http.StatusBadRequest,
			ErrorResponse{Description: "The receipt is invalid"})
		return
	}

	if err := ValidateReceipt(receipt); err != nil {
		WriteResponse(w, http.StatusBadRequest,
			ErrorResponse{Description: "The receipt is invalid"})
		return
	}

	points, err := CalculateReceiptPoints(receipt)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest,
			ErrorResponse{Description: "The receipt is invalid"})
		return
	}

	dbReceipt := db.Receipt{
		Points: points,
	}
	receiptID, err := db.CreateReceipt(dbReceipt)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest,
			ErrorResponse{Description: "The receipt is invalid"})
		return
	}

	response := ReceiptID{
		ID: receiptID,
	}
	WriteResponse(w, http.StatusOK, response)
}

type ReceiptPoints struct {
	Points int `json:"points"`
}

// receipts/:id/points
func ReceiptPointsHandler(w http.ResponseWriter, r *http.Request) {

	// TODO: Verify GET method
	// TODO: Load receipt poinyd by ID from database
	//	TODO: Return 404 if no receipt exists
	// TODO: Write receipt points to response and return 200

}
