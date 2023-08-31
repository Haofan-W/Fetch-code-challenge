package main

type ReceiptInput struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

type ReceiptOutput struct {
	ID     string `json:"id"`
	Points int    `json:"points"`
}

type Receipt struct {
	ReceiptInput
	ReceiptOutput
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type ReceiptWithPoint struct {
	Receipt Receipt
	Points  int
}

type PointsResponse struct {
	Points int
}
