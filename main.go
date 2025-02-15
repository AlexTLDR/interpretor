package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Messages []Message `json:"messages"`
	Model    string    `json:"model"`
}

type ChatResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("GROQ_API_KEY")

	request := ChatRequest{
		Messages: []Message{
			{
				Role:    "system",
				Content: "You are a direct translator. Provide only the translated words without any explanation or additional context.",
			},
			{
				Role:    "user",
				Content: "Translate 'add here your request' to German",
			},
		},
		Model: "mixtral-8x7b-32768",
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", "https://api.groq.com/openai/v1/chat/completions", bytes.NewBuffer(jsonData))

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Print raw response for debugging
	var rawResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&rawResponse); err != nil {
		log.Fatal(err)
	}

	prettyJSON, _ := json.MarshalIndent(rawResponse, "", "  ")
	fmt.Printf("Full API Response:\n%s\n", string(prettyJSON))
}
