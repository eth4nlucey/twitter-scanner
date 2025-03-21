package scraper

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ParseProfile extracts username, bio, and recent tweets from the HTML content
func ParseProfile(html string) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)
	}

	// Extract username
	username := doc.Find("div[data-testid='UserName'] span").First().Text()
	if username == "" {
		username = doc.Find("title").Text() // Backup: Page title
	}
	fmt.Println("Username:", username)

	// Extract bio
	bio := doc.Find("div[data-testid='UserDescription']").Text()
	if bio == "" {
		bio = doc.Find("div[data-testid='UserProfileHeader_Items']").Children().Eq(0).Text() // Ensures we get the actual bio, not "Joined" text
	}
	if bio == "" {
		bio = doc.Find("section[role='region'] div[dir='auto']").First().Text() // Last fallback
	}
	fmt.Println("Bio:", bio)

	// Extract tweets (Try different tweet containers)
	fmt.Println("\nRecent tweets:")
	doc.Find("div[data-testid='tweetText']").Each(func(i int, s *goquery.Selection) {
		fmt.Printf("- %s\n", s.Text())
	})
}
