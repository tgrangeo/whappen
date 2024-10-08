package main

import (
	"fmt"
	"io"
	"net/http"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
	"github.com/joho/godotenv"
	"github.com/tgrangeo/whappen/openAi"
	"github.com/tgrangeo/whappen/rss"
)

var qs = []*survey.Question{
	{
		Name: "menu",
		Prompt: &survey.Select{
			Message: "Welcome to Whappen:",
			VimMode: true,
			Options: []string{
				"📰 What's new",
				"💾 What's old/archived",
				"💡 Manage rss flux",
				"👋 quit"},
		},
	},
}

func listArticle(articles []rss.Article) {
	for {
		var titles []string
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
			return
		}
		openArticleMenu(selectedArticle)
	}
}

func openArticleMenu(art rss.Article) {
	articleMenu := []*survey.Question{
		{
			Name: "Article Menu",
			Prompt: &survey.Select{
				Message: art.Title + ":",
				Options: []string{
					"open in browser",
					"resume",
					"save for later",
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
			fmt.Println("WIP")
		case "👋 return":
			return
		}
	}

}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}
	for {
		answers := struct {
			Menu string
		}{}
		err := survey.Ask(qs, &answers)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		switch answers.Menu {
		case "📰 What's new":
			res := rss.FetchRSS()
			listArticle(res)
		case "💾 What's old/archived":
			fmt.Println("WIP")
		case "💡 Manage rss flux":
			fmt.Println("WIP")
		case "👋 quit":
			fmt.Println("see you next time 😄")
			return
		}
	}
}
