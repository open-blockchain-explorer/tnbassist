//go:build mock
// +build mock

package backup

import (
	"fmt"

	"github.com/open-blockchain-explorer/tnbassist/model"
)

// TNBAccounts is a helper fuction which backup accounts to a JSON file
func TNBAccounts(filename string, outdir string, accounts *model.Accounts) error {
	if len(*accounts) > 0 {
		return nil
	} else {
		return fmt.Errorf("No accounts")
	}
}
