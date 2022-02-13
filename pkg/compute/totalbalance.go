package compute

import "github.com/open-blockchain-explorer/tnbassist/model"

// AccountsTotalBalance computes the total balance of all accounts
func AccountsTotalBalance(accounts *model.Accounts) uint64 {
	var sum uint64
	for _, accountInfo := range *accounts {
		sum += accountInfo.Balance
	}
	return sum
}
