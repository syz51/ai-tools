package handlers

import "ai-web/internal/config"

type Handler struct {
	Config *config.Config
}

type TranslateResponse struct {
	TranslatedText string `json:"translatedText"`
}
