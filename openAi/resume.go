package openAi

import (
	"context"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)


func Resume(content []byte) {
	ctx := context.Background()
	key := os.Getenv("GPT_KEY")
	client := openai.NewClient(key)
	prompt := fmt.Sprintf("Please summarize this article: \n\n%s", content)
	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}
	resp, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		fmt.Println(err)
	}
	if len(resp.Choices) > 0 {
		fmt.Println("Summary:", resp.Choices[0].Message.Content)
	} else {
		fmt.Println("No response from OpenAI API")
	}
}
