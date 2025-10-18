package dto

type OllamaRequest struct {
	Model   string `json:"model"`
	Prompt  string `json:"prompt"`
	Stream  bool   `json:"stream"`
	Timeout int    `json:"timeout"`
}

type OllamaResponse struct {
	Response string `json:"response"`
}
