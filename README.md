# Receipt Processor

This is a receipt processing web service built in Go. It calculates points based on the
rules defined in the challenge and provides endpoints to submit receipts and retrieve
earned points.

## Prerequisites

Assuming you have already installed Go(https://golang.org/doc/install)

Note: Run `go version` to make sure your computer have go installed.
If your are sure you have go install, but 'go' command not found, try
modify your GO path on your device.

## Getting Started

1. Clone the repository:

   `git clone https://github.com/Haofan-W/Fetch-code-challenge.git`

   `cd Fetch-code-challenge`

### Mac

make sure your mac version is 1.13+ and have xCode and git installed.

### windows

make sure the git bash is installed.

2. Run the application with Go:

   `go run main.go structs.go receipt_calculator.go`

## Using the API

### Process Receipts

Endpoint: 'POST /receipts/process'

Submit a JSON receipt for processing. The API will return a JSON object containing an ID for the receipt.

### Example command on mac terminal:

```
curl -X POST -H "Content-Type: application/json" -d '{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    },
    {
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    },
    {
      "shortDescription": "Knorr Creamy Chicken",
      "price": "1.26"
    },
    {
      "shortDescription": "Doritos Nacho Cheese",
      "price": "3.35"
    },
    {
      "shortDescription": "Klarbrunn 12-PK 12 FL OZ",
      "price": "12.00"
    }
  ],
  "total": "35.35"
}' http://localhost:8080/receipts/process
```

### Example command on git bash:

```
curl -X POST -H "Content-Type: application/json" -d '{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    },
    {
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    },
    {
      "shortDescription": "Knorr Creamy Chicken",
      "price": "1.26"
    },
    {
      "shortDescription": "Doritos Nacho Cheese",
      "price": "3.35"
    },
    {
      "shortDescription": "Klarbrunn 12-PK 12 FL OZ",
      "price": "12.00"
    }
  ],
  "total": "35.35"
}' 'http://localhost:8080/receipts/process'
```

### Get Points

Endpoint: 'GET /receipts/{id}/points'

Retrieve the points for a receipt by ID

Example commmand:

`curl http://localhost:8080/receipts/your-receipt-id/points`

Note: your-receipt-id has to be the id we get from process receipts and the total points should be 28.
