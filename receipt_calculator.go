package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

func countAlphanumeric(s string) int {
	count := 0

	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			count++
		}
	}

	return count
}

func isRoundDollar(total string) bool {
	amount, err := strconv.ParseFloat(total, 64)
	if err != nil {
		return false
	}

	cents := amount - float64(int(amount))
	return cents == 0.0
}

func isMultipleOfQuarter(total string) bool {
	amount, err := strconv.ParseFloat(total, 64)
	if err != nil {
		return false
	}

	return math.Mod(amount, 0.25) == 0
}

func calculateItemTotalPoints(items []Item) int {
	totalPoints := 0

	for _, item := range items {
		itemPoints := calculateItemPoints(item)
		totalPoints += itemPoints
	}

	return totalPoints
}

func calculateItemPoints(item Item) int {
	trimmedDescription := strings.TrimSpace(item.ShortDescription)
	trimmedLength := len(trimmedDescription)

	if trimmedLength%3 == 0 {
		price, err := strconv.ParseFloat(item.Price, 64)
		if err != nil {
			return 0
		}
		points := int(math.Ceil(price * 0.2))
		return points
	}

	return 0
}

func calculatePurchaseDatePoint(purchaseDate string) int {
	dateParts := strings.Split(purchaseDate, "-")
	if len(dateParts) != 3 {
		fmt.Println("invaild date")
		return 0
	}

	day, err := strconv.Atoi(dateParts[2])
	if err != nil {
		return 0
	}

	if day%2 != 0 {
		return 6
	}

	return 0
}

func calculatePurchaseTimePoint(purchaseTime string) int {
	timeParts := strings.Split(purchaseTime, ":")
	if len(timeParts) != 2 {
		return 0
	}

	hour, err := strconv.Atoi(timeParts[0])
	if err != nil {
		return 0
	}

	if hour >= 14 && hour < 16 {
		return 10
	}

	return 0
}

// calculates the total points for the receipt
func calculateReceiptPoints(receipt Receipt) (int, error) {
	// One point for every alphanumeric character in the retailer name.
	totalPoints := countAlphanumeric(receipt.Retailer)

	// 50 points if the total is a round dollar amount with no cents.
	if isRoundDollar(receipt.Total) {
		totalPoints += 50
	}

	// 25 points if the total is a multiple of 0.25.
	if isMultipleOfQuarter(receipt.Total) {
		totalPoints += 25
	}

	// 5 points for every two items on the receipt.
	totalPoints += 5 * (len(receipt.Items) / 2)

	// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and
	// round up to the nearest integer. The result is the number of points earned.
	totalPoints += calculateItemTotalPoints(receipt.Items)

	// 6 points if the day in the purchase date is odd.
	totalPoints += calculatePurchaseDatePoint(receipt.PurchaseDate)

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	totalPoints += calculatePurchaseTimePoint(receipt.PurchaseTime)

	return totalPoints, nil
}
