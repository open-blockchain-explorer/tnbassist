//go:build !mock
// +build !mock

package backup

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/open-blockchain-explorer/tnbassist/model"
)

// TNBAccounts is a helper fuction which backup accounts to a JSON file
func TNBAccounts(filename string, outdir string, accounts *model.Accounts) error {
	jsonData, err := json.MarshalIndent(accounts, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fmt.Sprintf("%s/%s.json", outdir, filename), jsonData, os.ModePerm)
}
