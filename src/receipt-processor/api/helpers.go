package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func ReadJSONFromRequest(r *http.Request, obj interface{}) error {
	fn := "ReadJSONFromRequest"
	defer r.Body.Close()
	var bodyBytes []byte
	if _, err := r.Body.Read(bodyBytes); err != nil {
		return errors.New(fn + ": request body read failed - " + err.Error())
	}
	if err := json.Unmarshal(bodyBytes, obj); err != nil {
		return errors.New(fn + ": json unmarshal failed - " + err.Error())
	}
	return nil
}

func WriteResponse(w http.ResponseWriter, statusCode int, body interface{}) {
	fn := "WriteResponse"
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Println(fn+": marshal failed", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(bodyBytes); err != nil {
		log.Println(fn+": write failed", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println(string(bodyBytes))
	w.WriteHeader(statusCode)
}

// Receipts

func ValidateReceipt(receipt Receipt) error {
	// TODO: Validate receipt name is valid for regex
	// TODO: Validate date is correct format
	// TODO: Validate time
	// TODO: Validate total is valid for regex
	// TOOD: Loop over line items
	//	TODO: Validate description regex
	//	TODO: Validate prices regex
}

func CalculateReceiptPoints(receipt Receipt) (int, error) {
	// TODO: Calculate receipt points based on spec
}
