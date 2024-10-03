package rss

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mmcdole/gofeed"
)

func FetchRSS() {
	file, err := os.Open("./rss/rss_feeds.txt")
	if err != nil {
		log.Fatalf("Error opening the RSS feed list file: %v", err)
	}
	defer file.Close()
	fp := gofeed.NewParser()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := strings.TrimSpace(scanner.Text())
		if url == "" {
			continue
		}
		feed, err := fp.ParseURL(url)
		if err != nil {
			log.Printf("Error fetching the RSS feed from %s: %v", url, err)
			continue
		}
		fmt.Println("Feed Title:", feed.Title)
		for _, item := range feed.Items {
			fmt.Println("Title:", item.Title)
			fmt.Println("Link:", item.Link)
			fmt.Println("Published Date:", item.Published)
			fmt.Println("---\n\n")
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading the RSS feed list file: %v", err)
	}
}
