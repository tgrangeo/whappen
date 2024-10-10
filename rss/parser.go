package rss

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/mmcdole/gofeed"
)

type Article struct {
	Title string
	Link  string
	Date  string
	ToRead bool
}

func NewArticle(title, link, date string) Article {
	return Article{
		Title:    title,
		Link:     link,
		Date:     date,
		ToRead:   false,
	}
}

func FetchRSS() []Article {
	file, err := os.Open("./rss/rss_feeds.txt")
	if err != nil {
		log.Fatalf("Error opening the RSS feed list file: %v", err)
	}
	defer file.Close()
	fp := gofeed.NewParser()
	scanner := bufio.NewScanner(file)

	articleArray := []Article{}
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
		for _, item := range feed.Items {
			articleArray = append(articleArray, NewArticle(item.Title,item.Link,item.Published))
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading the RSS feed list file: %v", err)
	}
	return articleArray
}
