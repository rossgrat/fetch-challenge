package api

import (
	"log"
	"net/http"
)

type ReceiptProcessRequest struct {
}
type ReceiptProcessResponse struct {
}

// /receipts/process
func ReceiptsProcessHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("receipts/process")

	// TOOD: Verify POST method
	// TODO: Read request into JSON
	// TODO: Validate all JSON fields
	//	TOOD: Upon Failure, return 400
	// TODO: Save to database, database returns UUID
	// TODO: Calculate points for receipt, save under UUID
	// TODO: Write UUID to response and return 200
}

type ReceiptsPointsResponse struct {
}

// receipts/:id/points
func ReceiptPointsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("receipts/points")

	// TODO: Verify GET method
	// TODO: Load receipt poinyd by ID from database
	//	TODO: Return 404 if no receipt exists
	// TODO: Write receipt points to response and return 200

}
