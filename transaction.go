package monnify

import (
	"fmt"
	"net/http"
)

type Customer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Transactions struct {
	Content []Transaction `json:"content"`
}

type Transaction struct {
	TransactionReference string   `json:"transactionReference"`
	PaymentReference     string   `json:"paymentReference"`
	AmountPaid           string   `json:"amountPaid"`
	TotalPayable         string   `json:"totalPayable"`
	SettlementAmount     string   `json:"settlementAmount"`
	PaymentStatus        string   `json:"paymentStatus"`
	PaymentDescription   string   `json:"paymentDescription"`
	Currency             string   `json:"currency"`
	PaymentMethod        string   `json:"paymentMethod"`
	Customer             Customer `json:"email"`
}

type InitializeTransactionOption struct {
	Amount             float64  `json:"amount"`
	CustomerName       string   `json:"customerName"`
	CustomerEmail      string   `json:"customerEmail"`
	PaymentReference   string   `json:"paymentReference"`
	PaymentDescription string   `json:"paymentDescription"`
	CurrencyCode       string   `json:"currencyCode"`
	ContractCode       string   `json:"contractCode"`
	RedirectURL        string   `json:"redirectUrl"`
	PaymentMethods     []string `json:"paymentMethods"`
	Metadata           Metadata `json:"metadata,omitempty"`
}

type InitializedTransaction struct {
	TransactionReference string   `json:"transactionReference"`
	PaymentReference     string   `json:"paymentReference"`
	MerchantName         string   `json:"merchantName"`
	EnabledPaymentMethod []string `json:"enabledPaymentMethod"`
	CheckoutURL          string   `json:"checkoutUrl"`
}

// InitializeTransaction initializes the transaction that would be used for card payments and dynamic transfers
func (c *Client) InitializeTransaction(opts *InitializeTransactionOption) (*InitializedTransaction, error) {
	v := new(InitializedTransaction)

	_, err := c.doRequest(http.MethodPost,
		"/api/v1/merchant/transactions/init-transaction", opts, v)

	if err != nil {
		return v, err
	}

	return v, nil
}

// GetAllTransactions returns a list of transactions carried out on your integration.
func (c *Client) GetAllTransactions() ([]Transaction, error) {
	ts := new(Transactions)

	_, err := c.doRequest(http.MethodGet, "/api/v1/transactions/search", nil, ts)

	if err != nil {
		return nil, err
	}

	return ts.Content, nil
}

// GetTransactionStatus fetches the status of a transaction
func (c *Client) GetTransactionStatus(reference string) (*Transaction, error) {
	t := new(Transaction)

	_, err := c.doRequest(http.MethodGet,
		fmt.Sprintf("/api/v2/transactions/%s", reference), nil, t)

	if err != nil {
		return t, err
	}

	return t, nil
}
