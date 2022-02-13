package model

// Account represents a single account in the blockchain
type Account struct {
	AccountNumber string `json:"account_number" mapstructure:"account_number" validate:"required,len=64"`
	Balance       uint64 `json:"balance" mapstructure:"balance" validate:"required"`
	BalanceLock   string `json:"balance_lock" mapstructure:"balance_lock" validate:"required,len=64"`
}

// AccountInfo represents account information like balance and balance lock
type AccountInfo struct {
	Balance     uint64 `json:"balance" mapstructure:"balance" validate:"required"`
	BalanceLock string `json:"balance_lock" mapstructure:"balance_lock" validate:"required,len=64"`
}

// Accounts is a collection of Account
type Accounts map[string]AccountInfo

// PaginatedAccounts is a collection of Account with pagination
type PaginatedAccounts struct {
	Count    uint      `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Accounts []Account `json:"results"`
}
