package cmd

import (
	"fmt"
	"log"

	"github.com/eth4nlucey/twitter-scanner/pkg/scraper"
	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan [url]",
	Short: "Scans a Twitter profile for suspicious patterns",
	Long: `Scan a specified Twitter URL and analyze the content for suspicious patterns.

Example usage:
  twitter-scanner scan https://twitter.com/example`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		fmt.Println("Scanning URL:", url)

		content, err := scraper.ScrapeProfile(url)
		if err != nil {
			log.Fatalf("Error scraping: %v\n", err)
		}

		fmt.Println("Content retrieved:")
		fmt.Println(content)
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
