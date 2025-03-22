package Utils

import (
	"context"
	"encoding/json"
	"fmt"
	"genai2025/DTO"
	"log"
	"os"

	cohere "github.com/cohere-ai/cohere-go/v2"
	client "github.com/cohere-ai/cohere-go/v2/client"
)	

func RankResumes(apiKey string, baseProfile string, profiles []string) ([]string, error) {

	// payload := []string{baseProfile}
	// for _, u := range profiles {
	// 	payload = append(payload, u)
	// }

	// reqBody := cohereEmbedRequest{
	// 	Texts: payload,
	// 	Model: "embed-english-v3.0",
	// }

	// bodyBytes, err := json.Marshal(reqBody)
	// if err != nil {
	// 	return nil, err
	// }

	// // * Step 2 : send request to cohere
	// req, err := http.NewRequest(
	// 	"POST",
	// 	"https://api.cohere.ai/v1/embed",
	// )

	return nil, nil
}

func Test() {
	log.Print("Test")
}

type RankResult struct {
	Response *cohere.NonStreamedChatResponse
	Error    error
}

func RankCloseDevicesAsync(base_user DTO.UserPromptDTO, devices []DTO.UserPromptDTO) <-chan RankResult {
	resultChan := make(chan RankResult)
	go func() {
		apiKey := os.Getenv("COHERE_API_KEY")
		if apiKey == "" {
			resultChan <- RankResult{nil, fmt.Errorf("COHERE_API_KEY not set")}
			return
		}

		co := client.NewClient(client.WithToken(apiKey))
		resp, err := co.Chat(context.TODO(), &cohere.ChatRequest{
			Message: _gen_prompt(base_user, devices, 1, 0.6),
		})

		resultChan <- RankResult{resp, err}
	}()
	return resultChan
}

func _gen_prompt(user DTO.UserPromptDTO, profiles []DTO.UserPromptDTO, mode int, threshold float64) string {
	modeDescription := ""
	if mode == 1 {
		modeDescription = fmt.Sprintf("Only include profiles where the match score is **above %.2f**. These are considered best matches.", threshold)
	} else if mode == 2 {
		modeDescription = fmt.Sprintf("Only include profiles where the match score is **below %.2f**. These are considered worst matches.", threshold)
	}

	// ? Generate skills and interest 
	skillsJSON, err := json.Marshal(user.Skills)
	if err != nil {
		log.Fatal("Error marshaling skills:", err)
	}
	interestsJSON, err := json.Marshal(user.Interest)
	if err != nil {
		log.Fatal("Error marshaling interests:", err)
	}
	profileJSON, err := json.MarshalIndent(profiles, "", "\t")
	if err != nil {
		log.Fatal("error marshaling profiles:", err)
	}

	prompt := fmt.Sprintf(`
	You are helping me analyze networking matches based on a base profile and a list of other profiles.
	Your task is to compare each profile against the base profile in terms of skill overlap, complementary strengths, and potential for collaboration.
	Assign a match score from 0 to 1 based on these factors, where 1 is the highest possible match.
	%s

	Return the output strictly as a JSON object with the following structure. Do not change key names or format. Only change the data inside each field.

	{
	"ranking": [
		{
		"name": "",
		"ranking_position": 1,
		"skill_overlap": "",
		"complementary_strengths": "",
		"networking_potential": "",
		"suggested_collaboration": "",
		"reason": "",
		"match_score": 0.85
		}
	]
	}

	Here is the input:

	{
	"base": {
		"name": "%s",
		"skills": %s,
		"interests": %s
	},
	"profiles": %s
	}
	`, modeDescription, user.Username, skillsJSON, interestsJSON, profileJSON)

	return prompt
}

type cohereEmbedRequest struct {
	Texts []string `json:"texts"`
	Model string   `json:"model"`
}