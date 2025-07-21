package ai

import (
	"fmt"
)

// AIModel represents an AI model configuration.
type AIModel struct {
	Name    string
	APIKey  string
	Endpoint string
}

// Models stores the list of supported AI models.
var Models = []AIModel{
	{Name: "OpenAI GPT", APIKey: "", Endpoint: "https://api.openai.com/v1/completions"},
	{Name: "Google Bard", APIKey: "", Endpoint: "https://bard.google.com/api"},
}

// ConfigureAI allows configuring API keys for AI models.
func ConfigureAI(modelName, apiKey string) error {
	for i, model := range Models {
		if model.Name == modelName {
			Models[i].APIKey = apiKey
			fmt.Println("API Key configured for:", modelName)
			return nil
		}
	}
	return fmt.Errorf("AI model not found: %s", modelName)
}

// CallModel makes a request to the specified AI model.
func CallModel(modelName, prompt string) (string, error) {
	for _, model := range Models {
		if model.Name == modelName && model.APIKey != "" {
			// Mock response for now; integrate API calls in the future.
			return fmt.Sprintf("Response from %s: %s", modelName, prompt), nil
		}
	}
	return "", fmt.Errorf("API Key not configured for: %s", modelName)
}