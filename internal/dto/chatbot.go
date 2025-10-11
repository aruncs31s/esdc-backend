package dto

type ChatBotRequest struct {
	QueryMessage string `json:"query_message"`
}
type ChatBotResponse struct {
	Response string `json:"response"`
}
