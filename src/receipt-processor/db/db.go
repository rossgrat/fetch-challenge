package db

var uuidToReceipt map[string]int

type Receipt struct {
	Points int
}

func CreateReceipt(Receipt int) (string, error) {
	// TODO: Generate UUID for receipt
	// TODO: Save points for receipt
	return "", nil
}

func GetReceiptPoints(receiptID string) (int, error) {
	// TODO: Load points from map based on ID
	return -1, nil
}
