package cmd

import (
	"bytes"
	"encoding/json"
	"log"
	"time"

	"github.com/lestrrat-go/strftime"
	"github.com/open-blockchain-explorer/tnbassist/client"
	"github.com/open-blockchain-explorer/tnbassist/pkg/backup"
	"github.com/spf13/cobra"
)

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup <entity>",
	Short: "Backup data of the given <entity>",
	Long: `Backup command is used to fetch all the data given <entity>.
You can choose to either just the output as JSON or save to a JSON file. 
	
<entity> can be:
- accounts

For example:
tnbassist backup accounts -H 52.52.160.149 --save-as-json`,
	ValidArgs: []string{"accounts"},
	Args:      matchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		host, err := cmd.Flags().GetString("host")
		if err != nil {
			log.Fatal(err)
		}
		port, err := cmd.Flags().GetUint16("port")
		if err != nil {
			log.Fatal(err)
		}
		scheme, err := cmd.Flags().GetString("scheme")
		if err != nil {
			log.Fatal(err)
		}
		saveAsJSON, err := cmd.Flags().GetBool("save-as-json")
		if err != nil {
			log.Fatal(err)
		}
		strftimePattern, err := cmd.Flags().GetString("strftime")
		if err != nil {
			log.Fatal(err)
		}
		directory, err := cmd.Flags().GetString("outdir")
		if err != nil {
			log.Fatal(err)
		}

		thenewboston := client.NewTNBHTTPClient(
			client.HTTPConfig{
				Protocol: scheme,
				Host:     host,
				Port:     port,
			})
		accounts, err := thenewboston.FetchAllAccounts()
		if err != nil {
			log.Fatal(err)
		}

		if saveAsJSON {
			buf := new(bytes.Buffer)
			f, err := strftime.New(strftimePattern)
			if err != nil {
				log.Fatal(err)
			}
			if err := f.Format(buf, time.Now()); err != nil {
				log.Fatal(err)
			}
			if err := backup.TNBAccounts(buf.String(), directory, accounts); err != nil {
				log.Fatal(err)
			}
		} else {
			jsonData, err := json.MarshalIndent(accounts, "", "  ")
			if err != nil {
				log.Fatal(err)
			}
			cmd.Println(string(jsonData))
		}
	},
}

func init() {
	backupCmd.Flags().StringP("host", "H", "", "Base address/IP of primary validator server (required)")
	backupCmd.MarkFlagRequired("host")
	backupCmd.Flags().Uint16P("port", "p", 80, "Port of primary validator server")
	backupCmd.Flags().StringP("scheme", "s", "http", "URL Scheme/Protocol of primary validator server")
	backupCmd.Flags().StringP("strftime", "t", "%Y-%m-%d-%H_%M_%S", "String format time pattern used as filename for backup")
	backupCmd.Flags().StringP("outdir", "o", "./assets/account_backups", "Output directory path where backup will be saved")
	backupCmd.Flags().Bool("save-as-json", false, "Save to account backup in a JSON file")

	rootCmd.AddCommand(backupCmd)
}
