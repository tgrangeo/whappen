package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
	"github.com/joho/godotenv"
	database "github.com/tgrangeo/whappen/db"
	"github.com/tgrangeo/whappen/openAi"
	"github.com/tgrangeo/whappen/rss"
)

func listArticle(articles []rss.Article, db *sql.DB) {
	for {
		var titles []string
		if len(articles) == 0 {
			fmt.Println("\033[31mEmpty list\033[0m")
		}
		for _, article := range articles {
			titles = append(titles, article.Title)
		}
		titles = append(titles, "👋 Return")
		prompt := &survey.Select{
			Message:  "Choose an article:",
			PageSize: len(titles),
			Options:  titles,
		}
		var selectedTitle string
		err := survey.AskOne(prompt, &selectedTitle)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		var selectedArticle rss.Article
		for _, article := range articles {
			if article.Title == selectedTitle {
				selectedArticle = article
				break
			}
		}
		if selectedTitle == "👋 Return" {
			mainMenu(db)
		}
		openArticleMenu(selectedArticle, db)
	}
}

func openArticleMenu(art rss.Article, db *sql.DB) {
	laterValue := "save for later"
	if art.ToRead {
		laterValue = "mark as read"
	}
	articleMenu := []*survey.Question{
		{
			Name: "Article Menu",
			Prompt: &survey.Select{
				Message: art.Title + ":",
				Options: []string{
					"open in browser",
					"resume",
					laterValue,
					"👋 return"},
			},
		},
	}
	for {
		var response string
		err := survey.Ask(articleMenu, &response)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		switch response {
		case "open in browser":
			err = exec.Command("open", art.Link).Start()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
		case "resume":
			resp, err := http.Get(art.Link)
			if err != nil {
				fmt.Println("Error fetching article:", err)
				return
			}
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				fmt.Println("Error: Failed to fetch article, status code:", resp.StatusCode)
				return
			}
			rawContent, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Error reading article content:", err)
				return
			}
			openAi.Resume(rawContent)
		case "save for later":
			database.InsertToRead(db, art)
			return
		case "mark as read":
			database.RemoveArticle(db, art.Link)
			toReadList, err := database.FetchToReadArticle(db)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			listArticle(toReadList, db)
		case "👋 return":
			return
		}
	}

}

func mainMenu(db *sql.DB) {
	menu := []*survey.Question{
		{
			Name: "menu",
			Prompt: &survey.Select{
				Message: "Welcome to Whappen:",
				VimMode: true,
				Options: []string{
					"📰 What's new",
					"💾 What's to read",
					"💡 Manage rss flux",
					"👋 quit"},
			},
		},
	}
	for {
		answers := struct {
			Menu string
		}{}
		err := survey.Ask(menu, &answers)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		switch answers.Menu {
		case "📰 What's new":
			res := rss.FetchRSS()
			listArticle(res, db)
		case "💾 What's to read":
			toReadList, err := database.FetchToReadArticle(db)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			listArticle(toReadList, db)
		case "💡 Manage rss flux":
			fmt.Println("WIP")
		case "👋 quit":
			fmt.Println("see you next time 😄")
			os.Exit(0)
		}
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	mainMenu(db)
}
