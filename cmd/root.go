package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var rootCmd = &cobra.Command{
	Use:   "tnbassist",
	Short: "TNB Assist is CLI tool for thenewboston blockchain network",
	Long: `TNB Assist is a CLI (Command Line Interface) tool for thenewboston blockchain to perform various mundane tasks like taking daily accounts backup, computing statistics, etc easier.

Checkout our flagship blockchain explorer at https://tnbexplorer.com`,
	Run: func(cmd *cobra.Command, args []string) {
		// get docs flag
		generateDocs, err := cmd.Flags().GetBool("docs")
		if err != nil {
			log.Fatal(err)
		}
		// get version flag
		version, err := cmd.Flags().GetBool("version")
		if err != nil {
			log.Fatal(err)
		}
		// if version and docs flags is set then throw error
		if version && generateDocs {
			log.Fatal("--version and --docs flags are mutually exclusive")
			return
		}
		if generateDocs {
			createDirIfNotExist("./docs")
			err := doc.GenMarkdownTree(cmd, "./docs")
			if err != nil {
				log.Fatal(err)
			}
		} else if version {
			cmd.Println(cmd.Version)
		} else {
			cmd.Help()
		}
	},
}

func init() {
	// add generate docs flag
	rootCmd.Flags().Bool("docs", false, "Generate markdown docs for all commands")
	// add version flag
	rootCmd.Flags().BoolP("version", "v", false, "Print version information")
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

// matchAll is taken from https://github.com/spf12/cobra/issues/745#issuecomment-441326468
func matchAll(checks ...cobra.PositionalArgs) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		for _, check := range checks {
			if err := check(cmd, args); err != nil {
				return err
			}
		}
		return nil
	}
}
