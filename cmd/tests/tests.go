package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/rossgrat/fetch-challenge/src/receipt-processor/api"
)

const domain = "http://localhost:8080"

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
	receipt1 := noPointsReceipt
	receipt1.Retailer = "Target"

	resp1Status, resp1 := POST[api.ReceiptID](client, "/receipts/process", receipt1)
	if resp1Status != http.StatusOK {
		panic(errors.New("expected 200"))
	}
	resp2Status, resp2 := GET[api.ReceiptPoints](client, "/receipts/"+resp1.ID+"/points")
	if resp2Status != http.StatusOK {
		panic(errors.New("expected 200"))
	}
	if resp2.Points != 6 {
		panic(errors.New("expected 6 points"))
	}

	// TEST: Receipt round dollar amount and multiple of 0.25
	receipt2 := noPointsReceipt
	receipt2.Total = "1.00"

	resp3Status, resp3 := POST[api.ReceiptID](client, "/receipts/process", receipt2)
	if resp3Status != http.StatusOK {
		panic(errors.New("expected 200"))
	}
	resp4Status, resp4 := GET[api.ReceiptPoints](client, "/receipts/"+resp3.ID+"/points")
	if resp4Status != http.StatusOK {
		panic(errors.New("expected 200"))
	}
	if resp4.Points != 75 {
		panic(errors.New("expected 75 points"))
	}

	// TEST: Multiple of 0.25
	receipt3 := noPointsReceipt
	receipt3.Total = "1.25"

	resp5Status, resp5 := POST[api.ReceiptID](client, "/receipts/process", receipt3)
	if resp5Status != http.StatusOK {
		panic(errors.New("expected 200"))
	}
	resp6Status, resp6 := GET[api.ReceiptPoints](client, "/receipts/"+resp5.ID+"/points")
	if resp6Status != http.StatusOK {
		panic(errors.New("expected 200"))
	}
	if resp6.Points != 25 {
		panic(errors.New("expected 25 points"))
	}

}

// TODO: test with bad data
// TODO: test nonexsting endpoints
// TODO: test with wrong methods
// TODO: bad JSON
// TOOD: test receipts out of order

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
