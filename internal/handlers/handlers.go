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
		Config: cfg,
	}
}

func (h *Handler) Index(c echo.Context) error {
	return pages.Index().Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) Translate(c echo.Context) error {
	// Parse the incoming form data
	inputText := c.FormValue("inputText")
	targetLanguage := c.FormValue("targetLanguage")

	if inputText == "" || targetLanguage == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request: Empty Field")
	}

	// Construct the URL
	url, err := JoinUrl(h.Config.OpenaiServiceUrl, "translate")
	if err != nil {
		c.Logger().Errorf("Failed to parse and join URL: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	var data TranslateResponse
	resp, err := resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]string{"text": inputText, "targetLanguage": targetLanguage}).
		SetResult(&data).
		Post(url)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "HTTP request failed: "+err.Error())
	}
	if resp.StatusCode() != http.StatusOK {
		return echo.NewHTTPError(resp.StatusCode(), "Failed to translate")
	}

	// Send the HTML response
	return c.String(http.StatusOK, html.EscapeString(data.TranslatedText))
}
