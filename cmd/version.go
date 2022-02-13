package cmd

import (
	"github.com/spf13/cobra"
)

const version = "0.0.1-alpha"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Current version of TNB Assist CLI.",
	Long: `Current version of TNB Assist CLI in semantic versioning format.

Given a version number MAJOR.MINOR.PATCH, increment the:
- MAJOR version when you make incompatible API changes,
- MINOR version when you add functionality in a backwards compatible manner, and
- PATCH version when you make backwards compatible bug fixes.
- Additional labels for pre-release and build metadata are available as extensions to the MAJOR.MINOR.PATCH format.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println(rootCmd.Name(), "version", version)
	},
}

func init() {
	rootCmd.Version = version
	rootCmd.AddCommand(versionCmd)
}
