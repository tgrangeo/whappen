package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List articles from the RSS feeds",
	Long: `The 'list' command fetches and displays articles from your configured RSS feeds.
You can use flags to filter unread articles or articles from a specific source.
Examples:
  - List all articles:         list
  - List only unread articles: list --unread
  - List articles by source:   list --source="https://example.com/rss"`,
	Run: func(cmd *cobra.Command, args []string) {
		unread, err := cmd.Flags().GetBool("unread")
		if err != nil {
			fmt.Println("Error retrieving 'unread' flag:", err)
			return
		}
		source, err := cmd.Flags().GetString("source")
		if err != nil {
			fmt.Println("Error retrieving 'source' flag:", err)
			return
		}
		if unread {
			fmt.Println("Listing unread articles only...")
		} else {
			fmt.Println("Listing all articles...")
		}
		if source != "" {
			fmt.Printf("Filtering articles by source: %s\n", source)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("unread", "u", false, "List unread articles only")
	listCmd.Flags().StringP("source", "s", "", "List for a specific source")
}
