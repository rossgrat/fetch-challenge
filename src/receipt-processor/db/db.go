package db

// Per the spec, we are using a non-persisten form of data storage.
// Since the only thing we need to store is the receipt, and the
// receipt needs to be accessible by its ID, we have this private
// variable.
var idToReceipt map[string]any

type Receipt struct {
}

func CreateReceipt(receipt Receipt) string {

	// TODO: return receipt unique ID
	return ""
}

func GetReceipt(id string) Receipt {

	// TODO: Return receipt object for a map
	return Receipt{}

}
