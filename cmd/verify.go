package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/open-blockchain-explorer/tnbassist/model"
	"github.com/open-blockchain-explorer/tnbassist/pkg/compute"
	"github.com/spf13/cobra"
)

const (
	// MaxPointValue is the total coins in the network
	MaxPointValue = 281474976710656
)

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify <entity>",
	Short: "Verify data of the given <entity>.",
	Long: `Verify command is used to verify data of the given <entity>.
Exists with non-zero status code if any one of the validation fails.

<entity> can be:
- accounts
	Total sum of balance across accounts equals to MAX_POINT_VALUE (281474976710656)

For example:
tnbassist backup verify accounts -f ./assets/account_backups/2021-12-03-02_16_59.json`,
	ValidArgs: []string{"accounts"},
	Args:      matchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		file, err := cmd.Flags().GetString("file")
		if err != nil {
			log.Fatal(err)
		}

		content, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		var accounts model.Accounts
		err = json.Unmarshal(content, &accounts)
		if err != nil {
			log.Fatal(err)
		}
		if totalBalance := compute.AccountsTotalBalance(&accounts); totalBalance == MaxPointValue {
			fmt.Println("Valid")
		} else {
			fmt.Println("Invalid")
			os.Exit(1)
		}
	},
}

func init() {
	verifyCmd.Flags().StringP("file", "f", "", "Path to account backup JSON file (required)")
	verifyCmd.MarkFlagRequired("file")

	backupCmd.AddCommand(verifyCmd)
}
