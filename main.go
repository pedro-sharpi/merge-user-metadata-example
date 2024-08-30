package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the URL from the environment variable
	url := os.Getenv("CLERK_REQUEST_URL")
	token := os.Getenv("CLERK_TOKEN")

	if url == "" {
		log.Fatal("CLERK_REQUEST_URL environment variable is not set")
	}
	if token == "" {
		log.Fatal("CLERK_TOKEN environment variable is not set")
	}

	// Prepare the request body
	requestBody := map[string]interface{}{
		"public_metadata": map[string]interface{}{
			"customizations": map[string]interface{}{
				"enable_add_with_last_price": false,
			},
		},
		"private_metadata": map[string]interface{}{},
		"unsafe_metadata":  map[string]interface{}{},
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	// Create a new request
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Add the Authorization header
	req.Header.Add("Authorization", "Bearer "+token)

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making the request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading the response body: %v", err)
	}

	// Print the response status and body
	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("Body: %s\n", body)
}
