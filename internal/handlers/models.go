package handlers

import (
	"ai-web/internal/config"

	"github.com/go-resty/resty/v2"
)

type Handler struct {
	config       *config.Config
	openaiClient *resty.Client
}

type TranslateResponse struct {
	TranslatedText string `json:"output"`
}

type TranslateRequest struct {
	Input struct {
		Text           string `json:"text"`
		TargetLanguage string `json:"target_language"`
	} `json:"input"`
}
