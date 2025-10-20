
package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"strconv"

	"go-gin-app/config"
)

type OpenAIRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
}

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIResponse struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Message ChatMessage `json:"message"`
}

func GetOpenAIPromptFission(prompt string) ([]string, error) {
	template := config.Cfg.PromptFissionTemplate

	// Format the template with the prompt and fission count
	formattedPrompt := strings.Replace(template, "{prompt}", prompt, -1)
	countStr := strconv.Itoa(config.Cfg.PromptFissionCount)
	formattedPrompt = strings.Replace(formattedPrompt, "{prompt_fission_count}", countStr, -1)

	requestBody := OpenAIRequest{
		Model: "gpt-3.5-turbo",
		Messages: []ChatMessage{
			{
				Role:    "user",
				Content: formattedPrompt,
			},
		},
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", config.Cfg.OpenAIAPIUrl+"/v1/chat/completions", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.Cfg.OpenAIAPIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var openAIResp OpenAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&openAIResp); err != nil {
		return nil, err
	}

	if len(openAIResp.Choices) > 0 {
		content := openAIResp.Choices[0].Message.Content
		return strings.Split(content, "\n"), nil
	}

	return []string{}, nil
}
