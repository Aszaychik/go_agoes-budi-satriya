package service

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

type RecommendationService interface {
	GetRecommendations(request string) (string, error)
}

type RecommendationServiceImpl struct {
	OpenAPIKey string
}

func NewRecommendationService(OpenAPIKey string) RecommendationService {
	return &RecommendationServiceImpl{OpenAPIKey: OpenAPIKey}
}

func (service *RecommendationServiceImpl) GetRecommendations(request string) (string, error) {
	ctx := context.Background()
	openAPIKey := service.OpenAPIKey
	client := openai.NewClient(openAPIKey)

	resp, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleSystem, Content: "You are a helpful assistant that recommends laptops."},
			{Role: openai.ChatMessageRoleUser, Content: request},
		},
	})

	if err != nil {
		return "", fmt.Errorf("ChatCompletion error: %v", err)
	}

	return resp.Choices[0].Message.Content, nil
}