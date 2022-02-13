package cmd

import (
	"github.com/spf13/cobra"
)

// statsCmd represents the stats command
var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Stats command",
	Long:  `Stats command is used to get verious stats of the blockchain like richlist, quantiles, timesheets, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
