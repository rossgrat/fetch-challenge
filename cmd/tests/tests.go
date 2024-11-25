package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/rossgrat/fetch-challenge/src/receipt-processor/api"
)

const domain = "http://localhost:8080"

// All tests are present here
func main() {
	client := http.Client{}

	noPointsReceipt := api.Receipt{
		Retailer:     "&",
		PurchaseDate: "2000-02-02",
		PurchaseTime: "12:00",
		Items:        []api.Item{},
		Total:        "1.01",
	}

	// TEST: Receipt name points are valid
	receipt := noPointsReceipt
	receipt.Retailer = "Target"
	TestProcessReceipt(client, receipt, 6)

	receipt = noPointsReceipt
	receipt.Retailer = "Target&"
	TestProcessReceipt(client, receipt, 6)

	// TEST: Receipt round dollar amount and multiple of 0.25
	receipt = noPointsReceipt
	receipt.Total = "1.00"
	TestProcessReceipt(client, receipt, 75)

	// TEST: Multiple of 0.25
	receipt = noPointsReceipt
	receipt.Total = "1.25"
	TestProcessReceipt(client, receipt, 25)

	// TEST: Two items on receipt, description length not multiple of 3
	receipt = noPointsReceipt
	receipt.Items = []api.Item{
		{
			ShortDescription: "Item",
			Price:            "4.00",
		},
		{
			ShortDescription: "Item",
			Price:            "4.00",
		},
	}
	TestProcessReceipt(client, receipt, 5)

	// TEST: Three items on receipt, description length not multiple of 3
	receipt.Items = append(receipt.Items, api.Item{
		ShortDescription: "Item",
		Price:            "4.00",
	})
	TestProcessReceipt(client, receipt, 5)

	// TEST: Trimmed length of item is multiple of 3
	receipt = noPointsReceipt
	receipt.Items = []api.Item{
		{
			ShortDescription: "aaa",
			Price:            "18.00",
		},
	}
	TestProcessReceipt(client, receipt, 4)

	// TEST: Day in purchase date is odd
	receipt = noPointsReceipt
	receipt.PurchaseDate = "2000-01-01"
	TestProcessReceipt(client, receipt, 6)

	receipt = noPointsReceipt
	receipt.PurchaseDate = "2000-01-02"
	TestProcessReceipt(client, receipt, 0)

	// TEST: Time is after 2 and before 4
	receipt = noPointsReceipt
	receipt.PurchaseTime = "14:01"
	TestProcessReceipt(client, receipt, 10)

	receipt = noPointsReceipt
	receipt.PurchaseTime = "15:59"
	TestProcessReceipt(client, receipt, 10)

	receipt = noPointsReceipt
	receipt.PurchaseTime = "16:00"
	TestProcessReceipt(client, receipt, 0)

	receipt = noPointsReceipt
	receipt.PurchaseTime = "14:00"
	TestProcessReceipt(client, receipt, 0)

	// TEST: Load simple receipt 1, verify matches provided value
	receiptBytes, err := os.ReadFile("data/readme-receipt.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(receiptBytes, &receipt); err != nil {
		panic(err)
	}
	TestProcessReceipt(client, receipt, 28)

	// TEST: Load simple receipt 2, verify matches provided value
	receiptBytes, err = os.ReadFile("data/readme-receipt-2.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(receiptBytes, &receipt); err != nil {
		panic(err)
	}
	TestProcessReceipt(client, receipt, 109)

}

// Make a request to process a receipt, then turn around and get the points
// for that receipt
func TestProcessReceipt(client http.Client, receipt api.Receipt, expectedPoints int) {
	resp1Status, resp1 := POST[api.ReceiptID](client, "/receipts/process", receipt)
	if resp1Status != http.StatusOK {
		fmt.Println("expected 200")
		os.Exit(1)
	}
	resp2Status, resp2 := GET[api.ReceiptPoints](client, "/receipts/"+resp1.ID+"/points")
	if resp2Status != http.StatusOK {
		fmt.Println("expected 200")
		os.Exit(1)
	}
	if resp2.Points != expectedPoints {
		fmt.Printf("expected %d points", expectedPoints)
		os.Exit(1)
	}
}

// Marshal provided data, POST it to the provided path,
// then unmarshal and return the output if response is OK
func POST[T any](client http.Client, path string, body interface{}) (int, T) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest(
		http.MethodPost,
		domain+path,
		bytes.NewReader(bodyBytes))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println(http.MethodPost, path, resp.Status)
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

	if resp.StatusCode != http.StatusOK {
		var nothing T
		return resp.StatusCode, nothing
	}

	var respBody T
	if err := json.Unmarshal(b, &respBody); err != nil {
		panic(err)
	}

	return resp.StatusCode, respBody
}

// Retrieve data from the provided path via GET, unmarshal data
// if response is OK
func GET[T any](client http.Client, path string) (int, T) {
	req, err := http.NewRequest(
		http.MethodGet,
		domain+path,
		bytes.NewReader([]byte{}))
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println(http.MethodPost, path, resp.Status)
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

	if resp.StatusCode != http.StatusOK {
		var nothing T
		return resp.StatusCode, nothing
	}

	var respBody T
	if err := json.Unmarshal(b, &respBody); err != nil {
		panic(err)
	}

	return resp.StatusCode, respBody
}
