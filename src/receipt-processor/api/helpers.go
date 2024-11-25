package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/rossgrat/fetch-challenge/src/logger"
)

// Given an HTTP response write, a status code, and a body, perform the necessary
// marshalling and and write the status code and body to the response
func WriteResponse(w http.ResponseWriter, r *http.Request, statusCode int, body interface{}) {
	fn := "WriteResponse"
	bodyBytes, err := json.Marshal(body)
	w.WriteHeader(statusCode)
	if err != nil {
		logger.LogInfo(nil, fn+": marshal failed - "+err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(bodyBytes); err != nil {
		logger.LogInfo(nil, fn+": write failed -"+err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// Receipts

// Validate each field of the receipt. We use some common regex
// for short, descriptive fields with multiple words, and amounts
// that are formatted like USD.
// If any validation fails, return an error with the failing field
func ValidateReceipt(receipt Receipt) error {
	fn := "ValidateReceipt"
	wordsRegex, err := regexp.Compile("^[\\w\\s\\-&]+$")
	if err != nil {
		return errors.New(fmt.Sprintf("%s: failed to compile words regex - %s", fn, err.Error()))
	}
	amountRegex, err := regexp.Compile("^\\d+\\.\\d{2}$")
	if err != nil {
		return errors.New(fmt.Sprintf("%s: failed to compile amount regex - %s", fn, err.Error()))
	}

	if !wordsRegex.MatchString(receipt.Retailer) {
		return errors.New(fmt.Sprintf("%s: invalid retailer (%s)", fn, receipt.Retailer))
	}
	if _, err := time.Parse(time.DateOnly, receipt.PurchaseDate); err != nil {
		return errors.New(fmt.Sprintf("%s: invalid purchase date (%s) - %s", fn, receipt.PurchaseDate, err.Error()))
	}
	if _, err := time.Parse("15:04", receipt.PurchaseTime); err != nil {
		return errors.New(fmt.Sprintf("%s: invalid purchase time (%s) - %s", fn, receipt.PurchaseTime, err.Error()))
	}
	if !amountRegex.MatchString(receipt.Total) {
		return errors.New(fmt.Sprintf("%s: invalid total (%s)", fn, receipt.Total))
	}

	for _, item := range receipt.Items {
		if !wordsRegex.MatchString(item.ShortDescription) {
			return errors.New(fmt.Sprintf("%s: invalid item description (%s)", fn, item.ShortDescription))
		}
		if !amountRegex.MatchString(item.Price) {
			return errors.New(fmt.Sprintf("%s: invalid item amount (%s)", fn, item.Price))
		}

	}
	return nil
}

// We assume a valid receipt
func CalculateReceiptPoints(receipt Receipt) int {
	points := 0
	// One point for every alphanumeric character in the retailer name
	for _, char := range receipt.Retailer {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			points = points + 1
		}
	}

	// 50 points if the total is a round dollar amount with no cents
	totalFloat, _ := strconv.ParseFloat(receipt.Total, 64)
	if math.Mod(totalFloat, 1) == 0 {
		points = points + 50
	}

	// 25 points if the total is a multiple of 0.25
	if math.Mod(totalFloat, 0.25) == 0 {
		points = points + 25
	}

	// 5 points for every two items on the receipt
	points = points + (len(receipt.Items) / 2)

	// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned
	for _, item := range receipt.Items {
		trimmedDescription := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDescription)%3 == 0 {
			priceFloat, _ := strconv.ParseFloat(item.Price, 64)
			points = points + int(math.Ceil(priceFloat*0.2))
		}
	}

	// 6 points if the day in the purchase date is odd
	purchaseDate, _ := time.Parse(time.DateOnly, receipt.PurchaseDate)
	if purchaseDate.Day()%2 == 1 {
		points = points + 6
	}

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm
	purchaseTime, _ := time.Parse(time.TimeOnly, receipt.PurchaseTime)
	purchaseTimeHour, _, _ := purchaseTime.Clock()
	if purchaseTimeHour > 14 && purchaseTimeHour < 16 {
		points = points + 10
	}

	return points
}
