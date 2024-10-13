package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func main() {
	godotenv.Load("../.env")
	apiKey := os.Getenv("API_KEY")

	ctx := context.Background()

	client := openai.NewClient(apiKey)

	request := openai.CompletionRequest{
		Model:     openai.GPT3Babbage002,
		MaxTokens: 5,
		Prompt:    []string{"What size of shoes can fit a 12 year old?"},
	}

	resp, err := client.CreateCompletion(ctx, request)

	if err != nil {
		fmt.Printf("Error: %s \n", err)
	} else {
		fmt.Printf("Answer:\n %s \n", resp.Choices[0].Text)
	}

}
