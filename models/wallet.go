package models

type Wallet struct {
	Id          string                 `json:"id,omitempty"`
	Customer_id string                 `json:"customer_id,omitempty"`
	MetaData    map[string]interface{} `json:"meta_data,omitempty"`
	Balances    map[string]interface{} `json:"balances,omitempty"`
	CreatedAt   string                 `json:"created_at,omitempty"`
	UpdatedAt   string                 `json:"updated_at,omitempty"`
}

type Currencies struct {
	Currencies []Currency `json:"currencies"`
}

type Currency struct {
	Symbol       string `json:"symbol"`
	Name         string `json:"name"`
	SymbolNative string `json:"symbol_native"`
	Code         string `json:"code"`
	NamePlural   string `json:"name_plural"`
}

type Balances struct {
	Balances []Balance `json:"balances"`
}

type Balance struct {
	CurrencyCode   string  `json:"currency_code"`
	CurrencySymbol string  `json:"symbol"`
	Amount         float64 `json:"amount"`
	CreatedAt      string  `json:"created_at,omitempty"`
	UpdatedAt      string  `json:"updated_at,omitempty"`
}

type Transaction struct {
	Id          string `json:"id,omitempty"`
	Customer_id string `json:"customer_id,omitempty"`
	// Wallet_id    string  `json:"wallet_id"`
	Amount       float64 `json:"amount"`
	Type         string  `json:"type"`
	CurrencyCode string  `json:"currency_code"`
	Symbol       string  `json:"symbol"`
	Date         string  `json:"date"`
}
