package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/AlexTLDR/interpretor/internal/models"
)

type GroqClient struct {
	APIKey string
}

func NewGroqClient(apiKey string) *GroqClient {
	return &GroqClient{
		APIKey: apiKey,
	}
}

func (c *GroqClient) Translate(text string) (string, error) {
	request := models.ChatRequest{
		Messages: []models.Message{
			{
				Role:    "system",
				Content: "You are a direct translator. Provide only the translated words without any explanation or additional context.",
			},
			{
				Role:    "user",
				Content: text,
			},
		},
		Model: "mixtral-8x7b-32768",
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://api.groq.com/openai/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var response models.ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	if len(response.Choices) == 0 {
		return "", fmt.Errorf("no response choices received from the API")
	}

	content := response.Choices[0].Message.Content
	return strings.Trim(content, "'"), nil
}
