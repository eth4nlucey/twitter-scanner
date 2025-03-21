package cmd

import (
	"fmt"
	"twitter-scanner/pkg/scraper"

	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan [url]",
	Short: "Scans a Twitter profile",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		fmt.Println("\nScanning URL:", url)

		// Fetch profile with headless browser
		html := scraper.FetchProfile(url)

		// Parse the HTML
		scraper.ParseProfile(html)
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
