package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/AbhayAysola/tnb-go/account"
	"github.com/spf13/cobra"
)

// accountCmd represents the account command
var accountCmd = &cobra.Command{
	Use:       "account",
	Short:     "",
	Long:      `Returns an Account struct with a randomly generated keypair`,
	ValidArgs: []string{"create", "verify", "restore" /*"balance", "transactions"*/},
	Args:      matchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "create":
			acc, err := account.CreateAccount("")
			if err != nil {
				log.Fatal(err)
			}
			cmd.Println("Signing Key:\t", acc.SigningKeyHex)
			cmd.Println("Account No.:\t", acc.AccountNumberHex)
		case "restore":
			missingFlagNames := []string{}
			signingKey, err := cmd.Flags().GetString("sk")
			if err != nil {
				log.Fatal(err)
			}
			if signingKey == "" {
				missingFlagNames = append(missingFlagNames, "sk")
			}
			if len(missingFlagNames) > 0 {
				cmd.Usage()
				cmd.Println()
				cmd.Printf(`required flag(s) "%s" not set`, strings.Join(missingFlagNames, `", "`))
				return
			}
			acc, err := account.CreateAccount(signingKey)
			if err != nil {
				log.Fatal(err)
			}
			cmd.Println("Signing Key:\t", acc.SigningKeyHex)
			cmd.Println("Account No.:\t", acc.AccountNumberHex)
		case "verify":
			missingFlagNames := []string{}
			signingKey, err := cmd.Flags().GetString("sk")
			if err != nil {
				log.Fatal(err)
			}
			publicKey, err := cmd.Flags().GetString("acc")
			if err != nil {
				log.Fatal(err)
			}
			if signingKey == "" {
				missingFlagNames = append(missingFlagNames, "sk")
			}
			if publicKey == "" {
				missingFlagNames = append(missingFlagNames, "acc")
			}
			fmt.Println(signingKey, publicKey)
			if len(missingFlagNames) > 0 {
				cmd.Usage()
				cmd.Println()
				cmd.Printf(`required flag(s) "%s" not set`, strings.Join(missingFlagNames, `", "`))
				return
			}
			if err := cmd.ValidateArgs(args); err != nil {
				panic(err)
			}
			matched := account.VerifyKeyPair(signingKey, publicKey)
			if matched {
				cmd.Println("Valid")
			} else {
				cmd.Println("Invalid")
			}
		}
	},
}

func init() {
	accountCmd.Flags().StringP("sk", "s", "", "Signing key hex")
	accountCmd.Flags().StringP("acc", "a", "", "Account number hex")

	rootCmd.AddCommand(accountCmd)
}
