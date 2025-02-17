package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AlexTLDR/interpretor/internal/api"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("GROQ_API_KEY")
	groqClient := api.NewGroqClient(apiKey)

	result, err := groqClient.Translate(`Translate 
The Stuttgart Gophers are a local community of Go enthusiasts in Stuttgart, Germany. 
They cater to developers interested in the Go programming language (also known as Golang). 
The group organizes meetups, talks, and networking opportunities to share knowledge, 
learn from each other, and collaborate on projects. 
The Stuttgart Gophers are part of the global Go community and aim to promote the adoption and advancement of the language.

to German`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

}
