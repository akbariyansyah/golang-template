package rest

import (
	"net/http"
	"task_1/internal/pkg/formatter"

	"github.com/labstack/echo/v4"
)

type healthHandler struct{}

func NewHealthHandler(
	e *echo.Echo,
) *healthHandler {
	handler := &healthHandler{}
	v1 := e.Group("/v1")

	v1.GET("/ping", handler.health)
	return handler
}

func (h *healthHandler) health(c echo.Context) error {
	return c.JSON(http.StatusOK, formatter.NewSuccessResponse("pong"))
}
