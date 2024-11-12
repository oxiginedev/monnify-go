# Monnify Golang SDK

Simple wrapper for Monnify payment APIs


## Installation

```bash
go get github.com/oxiginedev/monnify-go
```

## Usage

```go
package main

import (
	"fmt"
	"log"

	"github.com/oxiginedev/monnify-go"
)

func main() {
    client, err := monnify.New(
        monnify.WithBaseURL("https://api.monnify.com"),
        monnify.WithAPIKey("your-api-key"),
        monnify.WithSecretKey("your-secret-key"),
    )

    if err != nil {
        log.Fatal(err)
    }

    transaction, err := client.InitializeTransaction(&monnify.InitializaTransactionOption{
        Amount:             300.50,
		CustomerName:       "Adedaramola Adetimehin",
		CustomerEmail:      "adedaramolaadetimehin@gmail.com",
		PaymentReference:   "skioenu3m3ovjo3v",
		PaymentDescription: "Subscription payment",

		CurrencyCode:       "NGN",
		ContractCode:       "5460134745",
		RedirectURL:        "https://redirect.url",
		PaymentMethods: []string{
			"CARD",
			"ACCOUNT_TRANSFER",
		},
    })

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(transaction.CheckoutURL)
}
```