package cmd

import (
	"fmt"
	"regexp"
	"time"

	"github.com/open-blockchain-explorer/tnbassist/client"
	"github.com/spf13/cobra"
)

// WIP:
// timesheetsCmd represents the timesheets command
var timesheetsCmd = &cobra.Command{
	Use:   "timesheets",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		github := client.NewGitHubGraphQLClient("")
		t := time.Date(2021, time.April, 15, 0, 0, 00, 0, time.UTC)
		// t = t.Truncate(24 * time.Hour)
		// t = t.Add(-24 * time.Hour)
		fmt.Println(t.Format(time.RFC3339))
		issues, _ := github.FetchNIssues("thenewboston-developers", "Contributor-Payments", client.Filters{
			Limit:      100,
			IssueState: client.Closed,
			FilterBy: client.IssueFilters{
				Since: &client.DateTime{
					Time: t,
				},
			},
			Labels: []string{"💰 Paid 💰"},
		})
		accountPattern := regexp.MustCompile(`[0-9a-fA-F]{64}`)
		hourPattern := regexp.MustCompile(`(?i)Total Time Spent(.+|\n)(\s|~)*(\d+|\d*(\.?\d+))\s*h`)
		hourBreakupPattern := regexp.MustCompile(`(?i)Time Spent(.+|\n)(\s|~)*(\d+|\d*(\.?\d+))\s*h`)
		bountyPattern := regexp.MustCompile(`(?i)Amount Requested(.+|\n)(\s|~)*\d+`)

		for _, issue := range issues.Edges {
			fmt.Println(issue.Node.Title)
			fmt.Println(issue.Node.ClosedAt)
			fmt.Println(issue.Node.BodyURL)
			accounts := accountPattern.FindAllString(issue.Node.Body, -1)
			duration := hourPattern.FindAllString(issue.Node.Body, -1)
			fmt.Println(accounts)
			if len(duration) != 0 {
				fmt.Println("CORE TEAM")
				fmt.Println(duration)
			} else {
				bounty := bountyPattern.FindAllString(issue.Node.Body, -1)
				if len(bounty) != 0 {
					fmt.Println("BOUNTY")
					fmt.Println(bounty)
				} else {
					timeSpent := hourBreakupPattern.FindAllString(issue.Node.Body, -1)
					if len(timeSpent) == 1 {
						fmt.Println("CORE TEAM")
						fmt.Println([]string{timeSpent[0]})
					} else {
						fmt.Println("N/A")
					}
				}
			}
			// fmt.Println(issue.Node.Body)
		}
	},
}

func init() {
	timesheetsCmd.Flags().StringP("token", "t", "", "GitHub personal access token")
	timesheetsCmd.Flags().StringSliceP("labels", "b", []string{}, "Labels to filter GitHub issues")
	// statsCmd.AddCommand(timesheetsCmd)
}
