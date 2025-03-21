package scraper

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

// FetchProfile uses ChromeDP to render JavaScript pages properly
func FetchProfile(url string) string {
	// Create ChromeDP context with options to prevent bot detection
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false), // Keep visible for debugging
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("ignore-certificate-errors", true),
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"),
	)

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	// Fetch HTML from X (Twitter) using ChromeDP
	var html string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(5*time.Second),     // Allow time for JS to load
		chromedp.OuterHTML("html", &html), // Extract entire HTML
	)
	if err != nil {
		log.Fatal("Failed to fetch profile:", err)
	}

	return html
}
