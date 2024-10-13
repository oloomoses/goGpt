package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

func main() {
	godotenv.Load()
	api_key := os.Getenv("API_KEY")

	if api_key == "" {
		log.Fatalln("Missing API KEY")
	}

	ctx := context.Background()

	client := openai.NewClient(api_key)

	for {
		fmt.Print("\n\n> ")
		reader := bufio.NewReader(os.Stdin)
		quiz, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}
		completeRequest(ctx, client, quiz)
	}
}

func makeRequest(question string) openai.CompletionRequest {
	return openai.CompletionRequest{
		Model:     openai.GPT3Babbage002,
		Prompt:    []string{question},
		MaxTokens: 5,
	}
}

func completeRequest(ctx context.Context, client *openai.Client, question string) {
	request := makeRequest(question)

	response, err := client.CreateCompletion(ctx, request)

	if err != nil {
		fmt.Printf("Error: %s \n", err)
	} else {
		fmt.Printf("Answer: %s \n", response.Choices[0].Text)
	}
}
