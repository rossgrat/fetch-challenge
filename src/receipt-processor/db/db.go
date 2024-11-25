package db

import (
	"errors"

	"github.com/google/uuid"
)

var uuidToReceipt map[uuid.UUID]Receipt

type Receipt struct {
	Points int
}

func CreateReceipt(receipt Receipt) (string, error) {
	id := uuid.New()
	uuidToReceipt[id] = receipt
	return id.String(), nil
}

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
