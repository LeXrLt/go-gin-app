
package utils

import (
	"fmt"
	"testing"
	"os"

	"go-gin-app/config"
)

func TestMain(m *testing.M) {
	// Setup: Load configuration
	if err := config.LoadConfig(); err != nil {
		fmt.Println("Failed to load config:", err)
		os.Exit(1)
	}
	// Run tests
	exitCode := m.Run()
	// Teardown
	os.Exit(exitCode)
}


func TestGetOpenAIPromptFission(t *testing.T) {
	// Skip this test if the API key and URL are not set,
	// as it requires actual API credentials.
	if config.Cfg.OpenAIAPIUrl == "your_openai_api_url" || config.Cfg.OpenAIAPIKey == "your_openai_api_key" {
		t.Skip("Skipping test: OpenAI API URL or Key not set in config.yaml")
	}

	prompt := "cat"
	fissionedPrompts, err := GetOpenAIPromptFission(prompt)

	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}

	if len(fissionedPrompts) == 0 {
		t.Fatal("Expected fissioned prompts, but got an empty list")
	}

	// You can add more assertions here, e.g., check the content of the prompts
	fmt.Printf("Got fissioned prompts: %v\n", fissionedPrompts)
}
