package monnify

type CreateReservedAccountGeneralOptions struct {
	AccountReference     string   `json:"accountReference"`
	AccountName          string   `json:"accountName"`
	CurrencyCode         string   `json:"currencyCode"`
	ContractCode         string   `json:"contractCode"`
	CustomerEmail        string   `json:"customerEmail"`
	CustomerName         string   `json:"customerName"`
	BVN                  string   `json:"bvn"`
	NIN                  string   `json:"nin"`
	GetAllAvailableBanks bool     `json:"getAllAvailableBanks"`
	PreferredBanks       []string `json:"preferredBanks"`
}

type VirtualAccount struct {
	BankName      string `json:"bankName"`
	BankCode      string `json:"bankCode"`
	AccountName   string `json:"accountName"`
	AccountNumber string `json:"accountNumber"`
}

func (c *Client) CreateReservedAccountGeneral() {}
