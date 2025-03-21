package scraper

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

// ScrapeProfile fetches page content from a URL with realistic user-agent
func ScrapeProfile(url string) (string, error) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) "+
			"AppleWebKit/537.36 (KHTML, like Gecko) "+
			"Chrome/123.0.0.0 Safari/537.36"),
		chromedp.Flag("headless", true),
	)

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, timeoutCancel := context.WithTimeout(ctx, 30*time.Second)
	defer timeoutCancel()

	ctx, ctxCancel := chromedp.NewContext(ctx)
	defer ctxCancel()

	var content string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(`body`, chromedp.ByQuery),
		chromedp.OuterHTML(`html`, &content),
	)

	if err != nil {
		log.Printf("Chromedp error: %v", err)
		return "", err
	}

	return content, nil
}
