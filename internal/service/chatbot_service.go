package service

import (
	aiproviders "esdc-backend/ai_providers"
	"esdc-backend/internal/dto"
	"esdc-backend/internal/model"
	"esdc-backend/internal/repository"
	"fmt"
	"strings"
	"time"
)

type ChatBotService interface {
	// SaveMessage(askedBy string, role string, category *string, content string, chatBotID int) error
	// GetMessagesByChatBotID(chatBotID int) ([]model.Message, error)
	// GetMessageByUserID(userID int) ([]model.Message, error)
	// GetMessageByCategory(category string) ([]model.Message, error)
	Ask(user *string, question string) (dto.ChatBotResponse, error)
}
type chatBotService struct {
	chatBotRepo repository.ChatBotRepository
	ollamaRepo  repository.OllamaRepository
	ai_provider aiproviders.Ollama
	userRepo    repository.UserRepository
}

func NewChatBotService(chatBotRepo repository.ChatBotRepository, ollamaRepo repository.OllamaRepository, userRepo repository.UserRepository) ChatBotService {
	ai_provider := aiproviders.Ollama{Model: "tinyllama"}
	return &chatBotService{
		chatBotRepo: chatBotRepo,
		ollamaRepo:  ollamaRepo,
		ai_provider: ai_provider,
		userRepo:    userRepo,
	}
}
func (s *chatBotService) Ask(user *string, question string) (dto.ChatBotResponse, error) {
	response, err := s.askOllama(question)
	if err != nil {
		return dto.ChatBotResponse{}, err
	}

	newMessage := model.ChatBotMessage{
		AskedBy:   s.getUser(user),
		Content:   question,
		CreatedAt: time.Now(),
		Category:  getCategory(question),
		Response:  &response,
	}
	newOllamaResponse := model.Ollama{
		ID:        0,
		ModelName: "tinyllama",
		Prompt:    question,
		Response:  response,
		AskedBy:   s.getUser(user),
	}
	if err := s.ollamaRepo.SaveChat(&newOllamaResponse); err != nil {
		return dto.ChatBotResponse{}, err
	}
	if err := s.chatBotRepo.SaveMessage(&newMessage); err != nil {
		return dto.ChatBotResponse{}, err
	}
	fmt.Println("Saved message:", response)
	return dto.ChatBotResponse{
		Response: response,
	}, nil
}
func (s *chatBotService) getUser(username *string) *int {
	if username == nil || *username == "anonymous" {
		defValue := 0
		return &defValue
	}
	userID, err := s.userRepo.FindUserIDByUsername(*username)
	if err != nil {
		defValue := 0
		return &defValue
	}
	return &userID
}
func (s *chatBotService) askOllama(question string) (string, error) {
	response, err := s.ai_provider.AskOllama(question)
	if err != nil {
		return "", err
	}
	// Implement more ai providers.
	return response, nil
}

func getCategory(message string) *string {
	knownCategory := []string{
		"project",
		"product",
		"blog",
		"user",
	}
	for _, value := range knownCategory {
		if strings.Contains(strings.ToLower(message), value) {
			return &value
		} else if strings.Contains(strings.ToLower(message), value+"s") {
			return &value
		}
	}
	defaultValue := "general"
	return &defaultValue
}
