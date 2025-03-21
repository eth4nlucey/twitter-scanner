package main

import (
	"fmt"
	"log"
	"os"

	"github.com/eth4nlucey/twitter-scanner/pkg/scraper"
)

func main() {
	if len(os.Args) < 3 || os.Args[1] != "scan" {
		fmt.Println("Usage: go run main.go scan <URL>")
		return
	}

	url := os.Args[2]
	fmt.Println("\nScanning URL:", url)

	// Fetch and parse profile using ChromeDP
	html := scraper.FetchProfile(url)

	// 🚨 DEBUG: Print first 1500 characters of HTML to check if we got real content
	fmt.Println("\n🚨 DEBUG: HTML Snippet 🚨\n", html[:1500])

	if html == "" {
		log.Fatal("Failed to fetch profile: empty response")
		return
	}

	// Parse the fetched profile HTML
	scraper.ParseProfile(html)
}
