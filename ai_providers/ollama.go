package aiproviders

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Ollama struct {
	Model string
}

type GenerateRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream,omitempty"`
}

type GenerateResponse struct {
	Response string `json:"response"`
	Done     bool   `json:"done"`
}

func (o *Ollama) AskOllama(question string) (string, error) {
	// Ensure Ollama server is running: ollama serve

	req := GenerateRequest{
		Model:  o.Model,
		Prompt: question,
		Stream: true,
	}
	fmt.Println("Ollama request:", req)
	jsonData, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	resp, err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to call Ollama API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("Ollama API error: %s", string(body))
	}

	// Read the streamed response
	var fullResponse string
	decoder := json.NewDecoder(resp.Body)
	for {
		var chunk GenerateResponse
		if err := decoder.Decode(&chunk); err != nil {
			if err == io.EOF {
				break
			}
			return "", fmt.Errorf("error decoding response: %w", err)
		}
		fullResponse += chunk.Response
		if chunk.Done {
			break
		}
	}

	fmt.Println("Ollama response:", fullResponse)
	return fullResponse, nil
}
