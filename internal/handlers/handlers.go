package handlers

import (
	"ai-web/internal/config"
	"ai-web/pkg/view/pages"
	"html"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
)

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{
		config:       cfg,
		openaiClient: resty.New().SetBaseURL(cfg.OpenaiServiceUrl).SetHeader("Content-Type", "application/json"),
	}
}

func (h *Handler) Index(c echo.Context) error {
	return pages.Index().Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) Translate(c echo.Context) error {
	return pages.Translate().Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) TranslateText(c echo.Context) error {
	// Parse the incoming form data
	inputText := c.FormValue("inputText")
	targetLanguage := c.FormValue("targetLanguage")

	// TODO: Replace with validation library
	if inputText == "" || targetLanguage == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request: Empty Field")
	}

	req := new(TranslateRequest)
	req.Input.Text = inputText
	req.Input.TargetLanguage = targetLanguage

	var data TranslateResponse
	resp, err := h.openaiClient.R().
		SetBody(req).
		SetResult(&data).
		Post("/translation/text/invoke")

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "HTTP request failed: "+err.Error())
	}
	if resp.StatusCode() != http.StatusOK {
		return echo.NewHTTPError(resp.StatusCode(), "Failed to translate")
	}

	// Send the HTML response
	return c.String(http.StatusOK, html.EscapeString(data.TranslatedText))
}
