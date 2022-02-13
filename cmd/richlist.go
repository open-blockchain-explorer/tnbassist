package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/open-blockchain-explorer/tnbassist/pkg/compute"
	"github.com/spf13/cobra"
)

// richlistCmd represents the richlist command
var richlistCmd = &cobra.Command{
	Use:   "richlist",
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
		maxRichListCount, err := cmd.Flags().GetUint("count")
		if err != nil {
			log.Fatal(err)
		}
		outputAsJSON, err := cmd.Flags().GetBool("json")
		if err != nil {
			log.Fatal(err)
		}
		blacklist, err := cmd.Flags().GetStringSlice("blacklist")
		if err != nil {
			log.Fatal(err)
		}
		blackset := map[string]struct{}{}
		for _, account := range blacklist {
			blackset[account] = struct{}{}
		}

		stats, err := compute.CoumputeQuantiles(file, &compute.Options{
			Blacklist:        blackset,
			Quantiles:        []uint{},
			MaxRichListCount: maxRichListCount,
		})
		if err != nil {
			log.Fatal(err)
		}

		if outputAsJSON {
			jsonData, err := json.MarshalIndent(stats.RichList, "", "  ")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(jsonData))
		} else {
			fmt.Println("RANK\tBALANCE\tACCOUNT")
			for idx, richlist := range stats.RichList {
				fmt.Printf("%d\t%d\t%s\n", idx+1, richlist.Balance, richlist.Account)
			}
		}
	},
}

func init() {
	richlistCmd.Flags().StringP("file", "f", "", "Path to account backup JSON file")
	richlistCmd.MarkFlagRequired("file")
	richlistCmd.Flags().UintP("count", "c", 1, "Path to account backup JSON file")
	richlistCmd.Flags().StringSliceP("blacklist", "b", []string{}, "Path to account backup JSON file")
	richlistCmd.Flags().Bool("json", false, "Path to account backup JSON file")

	statsCmd.AddCommand(richlistCmd)
}
