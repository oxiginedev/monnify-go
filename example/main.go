package main

import (
	"fmt"
	"log"

	"github.com/oxiginedev/monnify-go"
)

func main() {
	mc, err := monnify.New(
		monnify.WithBaseURL("https://sandbox.monnify.com"),
		monnify.WithAPIKey("MK_TEST_ZTRGGXUFKZ"),
		monnify.WithSecretKey("DR4N1SFAL3U8FSZQXDC6JK1M9FW9FSRG"),
	)

	if err != nil {
		log.Fatal(err)
	}

	// transaction, err := mc.InitializeTransaction(&monnify.InitializeTransactionOption{
	// Amount:             300.50,
	// CustomerName:       "Adedaramola Adetimehin",
	// CustomerEmail:      "adedaramolaadetimehin@gmail.com",
	// PaymentReference:   "skioenu3m3ovjo3v",
	// PaymentDescription: "Subscription payment",

	// CurrencyCode:       "NGN",
	// ContractCode:       "5460134745",
	// RedirectURL:        "https://redirect.url",
	// PaymentMethods: []string{
	// 	"CARD",
	// 	"ACCOUNT_TRANSFER",
	// },
	// })

	transaction, err := mc.GetAllTransactions()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", transaction)
}
