package middleware

import (
	"errors"
	"net/http"
	"task_1/internal/domain/user"
	"task_1/internal/pkg/formatter"

	"github.com/labstack/echo/v4"
)

func ErrorHandlingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Defer to recover from panics
		defer func() {
			if err := recover(); err != nil {
				c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"message": "Internal Server Error",
					"error":   err,
				})
			}
		}()

		err := next(c)

		if err != nil {
			code := http.StatusInternalServerError
			if he, ok := err.(*echo.HTTPError); ok {
				code = he.Code
			}

			if errors.Is(err, user.ErrEmailAlreadyExist) {
				code = http.StatusBadRequest
			}

			if errors.Is(err, user.ErrEmailEmpty) {
				code = http.StatusBadRequest
			}

			if errors.Is(err, user.ErrEmailNotFound) {
				code = http.StatusBadRequest
			}

			c.JSON(code, formatter.NewErrorResponse(err))
			return nil
		}

		return nil
	}
}
