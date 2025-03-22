package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type CohereRequest struct {
	Model     string `json:"model"`
	Prompt    string `json:"prompt"`
	MaxTokens int    `json:"max_tokens"`
}

type CohereResponse struct {
	Generations []struct {
		Text string `json:"text"`
	} `json:"generations"`
}

 
func main() {
	fmt.Println("🚀 Starting Cohere Test...")

	apiKey := os.Getenv("COHERE_API_KEY")
	if apiKey == "" {
		fmt.Println("❌ API key not found — make sure COHERE_API_KEY is set")
		return
	}

	prompt := `
	User A:
	- School: University of Toronto
	- Skills: Python, Machine Learning, Data Science

	Nearby users:

	- User B:
	- School: University of Toronto
	- Skills: Python, AI, Statistics

	- User C:
	- School: University of British Columbia
	- Skills: JavaScript, Design, React

	Which user should User A connect with and why?`

	payload := CohereRequest{
		Model:     "command", // or "command-nightly"
		Prompt:    prompt,
		MaxTokens: 100,
	}

	jsonPayload, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "https://api.cohere.ai/generate", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("❌ Request failed:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("❌ Failed to read response body:", err)
		return
	}

	fmt.Println("🌐 Raw response from Cohere:")
	fmt.Println(string(body))

	var cohereResp CohereResponse
	err = json.Unmarshal(body, &cohereResp)
	if err != nil {
		fmt.Println("❌ Failed to parse JSON:", err)
		return
	}

	fmt.Println("✅ AI Suggestion:")
	if len(cohereResp.Generations) > 0 {
		fmt.Println(cohereResp.Generations[0].Text)
	} else {
		fmt.Println("⚠️ No response received from Cohere.")
	}
}
