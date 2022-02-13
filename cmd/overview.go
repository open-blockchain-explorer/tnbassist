package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/open-blockchain-explorer/tnbassist/client"
	"github.com/open-blockchain-explorer/tnbassist/pkg/backup"
	"github.com/open-blockchain-explorer/tnbassist/pkg/compute"
	"github.com/spf13/cobra"
)

// overviewCmd represents the overview command
var overviewCmd = &cobra.Command{
	Use:   "overview",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		file, err := cmd.Flags().GetString("file")
		if err != nil {
			log.Fatal(err)
		}
		blacklist, err := cmd.Flags().GetStringSlice("blacklist")
		if err != nil {
			log.Fatal(err)
		}
		quantiles, err := cmd.Flags().GetUintSlice("quantiles")
		if err != nil {
			log.Fatal(err)
		}
		outputAsJSON, err := cmd.Flags().GetBool("json")
		if err != nil {
			log.Fatal(err)
		}
		directory, err := cmd.Flags().GetString("outdir")
		if err != nil {
			log.Fatal(err)
		}
		appendToCSV, err := cmd.Flags().GetBool("append-to-csv")
		if err != nil {
			log.Fatal(err)
		}
		postToServer, err := cmd.Flags().GetBool("post-to-server")
		if err != nil {
			log.Fatal(err)
		}
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
		token, err := cmd.Flags().GetString("token")
		if err != nil {
			log.Fatal(err)
		}
		blackset := map[string]struct{}{}
		for _, account := range blacklist {
			blackset[account] = struct{}{}
		}

		stats, err := compute.CoumputeQuantiles(file, &compute.Options{
			Blacklist:        blackset,
			Quantiles:        quantiles,
			MaxRichListCount: 0,
		})
		if err != nil {
			log.Fatal(err)
		}

		if postToServer {
			if outputAsJSON {
				log.Fatal("--post-to-server and --json cannot be used together")
			}
			tnbexplorer := client.NewTNBExplorerHTTPClient(
				client.HTTPConfig{
					Protocol: scheme,
					Host:     host,
					Port:     port,
				}, token)

			// filename will be used as date to maintain reconciliation work easier
			filename := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
			statusCode, body, err := tnbexplorer.PostStats(stats.ToLegacyStats(filename))
			fmt.Printf("TNBExplorer POST STATS API responsed with status code: %d\n\n", statusCode)
			if err != nil || statusCode != http.StatusCreated {
				log.Fatal(err, string(body))
			}
		}

		if appendToCSV {
			// filename will be used as date to maintain reconciliation work easier
			filename := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
			err := backup.AppendToCSV("stats", directory, stats.ToLegacyStats(filename))
			if err != nil {
				log.Fatal(err)
			}
		}

		if outputAsJSON {
			jsonData, err := json.MarshalIndent(stats, "", "  ")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(jsonData))
		} else {
			fmt.Println("TOTAL:\t\t", stats.Total)
			fmt.Println("SHIFT:\t\t", stats.Shift)
			fmt.Println("ACCOUNTS:\t", stats.NAccounts)
			fmt.Println("MAX_BALANCE:\t", stats.MaxBalance)
			fmt.Println("RICHEST:\t", stats.RichestAccount)
			fmt.Println("QUANTILE\tOWNERSHIP\tACCOUNTS\tWEALTH")
			for pct, quantile := range stats.Quanitiles {
				fmt.Printf("%d\t\t%.2f%%\t\t%d\t\t%d\n", pct, quantile.Ownership*100, quantile.NAccounts, quantile.Wealth)
			}
		}
	},
}

func init() {
	overviewCmd.Flags().StringP("file", "f", "", "Path to account backup JSON file")
	overviewCmd.MarkFlagRequired("file")
	overviewCmd.Flags().StringSliceP("blacklist", "b", []string{}, "Path to account backup JSON file")
	overviewCmd.Flags().UintSliceP("quantiles", "q", []uint{5, 10, 25, 50}, "Path to account backup JSON file")
	overviewCmd.Flags().StringP("token", "t", "", "Path to account backup JSON file")
	overviewCmd.Flags().StringP("outdir", "o", "./assets", "Output directory path where backup will be saved")
	overviewCmd.Flags().Bool("append-to-csv", false, "Path to account backup JSON file (make sure atleast 5, 10, 25 and 50 quantiles are provided)")
	overviewCmd.Flags().StringP("host", "H", "", "Base address/IP of primary validator server (required)")
	overviewCmd.Flags().Uint16P("port", "p", 80, "Port of primary validator server")
	overviewCmd.Flags().StringP("scheme", "s", "http", "URL Scheme/Protocol of primary validator server")
	overviewCmd.Flags().Bool("post-to-server", false, "Path to account backup JSON file (make sure TNBExplorer API token is provided)")
	overviewCmd.Flags().Bool("json", false, "Output should be in JSON format")

	statsCmd.AddCommand(overviewCmd)
}
