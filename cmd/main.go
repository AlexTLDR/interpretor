package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
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

	tmpl := template.Must(template.ParseFiles("web/templates/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			Languages []api.Language
		}{
			Languages: api.SupportedLanguages,
		}
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/translate", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// Log the form data
			log.Printf("From Language: %s", r.FormValue("fromLang"))
			log.Printf("To Language: %s", r.FormValue("toLang"))
			log.Printf("Style: %s", r.FormValue("style"))
			log.Printf("Text to translate: %s", r.FormValue("text"))

			var formData struct {
				FromLang string `json:"fromLang"`
				ToLang   string `json:"toLang"`
				Style    string `json:"style"`
				Text     string `json:"text"`
			}

			if err := json.NewDecoder(r.Body).Decode(&formData); err != nil {
				log.Printf("Error decoding request body: %v", err)
				http.Error(w, "Invalid request format", http.StatusBadRequest)
				return
			}

			log.Printf("JSON Data received: %+v", formData)

			result, err := groqClient.Translate(formData.Text, formData.FromLang, formData.ToLang, api.TranslationStyle(formData.Style))
			if err != nil {
				log.Printf("Translation error: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"translation": result})
		}
	})

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
