package endpoints

import "net/http"

type ReceiptProcessRequest struct {
}
type ReceiptProcessResponse struct {
}

// /receipts/process
func ReceiptsProcessHandler(req *http.Request) {

	// TODO: Verify POST

}

type ReceiptsPointsResponse struct {
}

// receipts/:id/points
func ReceiptPointsHandler() {

	// TODO: Verify GET

}
