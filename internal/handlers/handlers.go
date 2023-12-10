package handlers

import (
	"ai-web/pkg/view/pages"

	"github.com/labstack/echo/v4"
)

func IndexHandler(c echo.Context) error {
	return pages.Index().Render(c.Request().Context(), c.Response().Writer)
}
