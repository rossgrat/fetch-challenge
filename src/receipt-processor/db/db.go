package db

import (
	"errors"

	"github.com/google/uuid"
)

var uuidToReceipt map[uuid.UUID]Receipt = make(map[uuid.UUID]Receipt)

type Receipt struct {
	Points int
}

// Create a receipt in the "database" and generate
// a UUID for the new object, return that ID
func CreateReceipt(receipt Receipt) (string, error) {
	id := uuid.New()
	uuidToReceipt[id] = receipt
	return id.String(), nil
}

// Load a receipt based on a provided UUID, if
// the ID is invalid or no receipt is present for
// the ID, return an error
func GetReceipt(receiptID string) (Receipt, error) {
	fn := "GetReceipt"
	id, err := uuid.Parse(receiptID)
	if err != nil {
		return Receipt{}, errors.New(fn + ": invalid UUID")
	}
	receipt, ok := uuidToReceipt[id]
	if !ok {
		return Receipt{}, errors.New(fn + ": no ID for receipt")
	}
	return receipt, nil
}
