package rest

import (
	"net/http"
	"task_1/internal/app/param"
	"task_1/internal/app/service"
	"task_1/internal/pkg/formatter"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(
	e *echo.Echo,
	service service.UserService) *userHandler {
	handler := &userHandler{userService: service}
	v1 := e.Group("/v1")

	v1.POST("/user/create", handler.CreateUser)
	return handler
}

func (h *userHandler) CreateUser(c echo.Context) error {
	ctx := c.Request().Context()
	data := new(param.UserCreate)
	if err := c.Bind(data); err != nil {
		return err
	}

	res, err := h.userService.CreateUser(ctx, data)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, formatter.NewSuccessResponse(res))
}
