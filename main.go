package main

import (
	"fmt"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
	"github.com/tgrangeo/whappen/rss"
)

var qs = []*survey.Question{
	{
		Name: "menu",
		Prompt: &survey.Select{
			Message: "Welcome to Whappen:\n",
			VimMode: true,
			Options: []string{
				"📰 What's new\n",
				"💾 What's old/archived\n",
				"💡 Manage rss flux\n",
				"👋 quit\n"},
		},
	},
}

func listArticle(articles []rss.Article) {
	var titles []string
	for _, article := range articles {
		titles = append(titles, article.Title)
	}
	prompt := &survey.Select{
		Message: "Choose an article:",
		Options: titles,
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
	err = exec.Command("open", selectedArticle.Link).Start()
}

func main() {
	answers := struct {
		Menu string
	}{}
	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	switch answers.Menu {
	case "📰 What's new\n":
		res := rss.FetchRSS()
		listArticle(res)
	case "💾 What's old/archived\n":
		fmt.Println("WIP")
	case "💡 Manage rss flux\n":
		fmt.Println("WIP")
	case "👋 quit\n":
		fmt.Println("see you next time 😄")
		return
	}
}
