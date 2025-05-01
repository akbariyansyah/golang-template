package middleware

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// responseWriter wraps echo.ResponseWriter to capture the response body
type responseWriter struct {
	body *bytes.Buffer
	http.ResponseWriter
	statusCode int
}

// Write captures the response body
func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.body.Write(b) // Write to buffer for logging
	return rw.ResponseWriter.Write(b)
}

// WriteHeader captures the status code
func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func LoggerMiddleware(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()

			var requestBody string
			if req.Body != nil {
				bodyBytes, _ := io.ReadAll(req.Body)
				requestBody = string(bodyBytes)
				req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			}

			// Wrap response writer
			rw := &responseWriter{
				body:           new(bytes.Buffer),
				ResponseWriter: res.Writer,
				statusCode:     http.StatusOK,
			}
			res.Writer = rw

			// Start timer
			start := time.Now()

			// Process the request
			err := next(c)

			// Calculate latency
			latency := time.Since(start)

			// Log request and response
			logger.Info("HTTP Request",
				zap.String("method", req.Method),
				zap.String("url", req.URL.String()),
				zap.Int("status", rw.statusCode),
				zap.Duration("latency", latency),
				zap.String("request_body", requestBody),
				zap.String("response_body", rw.body.String()),
			)

			return err
		}
	}
}

func CreateLogger(logFilePath string) (*zap.Logger, error) {
	// Create a file writer
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	// Create a Zap core with a file writer
	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(file),
		zapcore.InfoLevel,
	)

	// Create the logger
	return zap.New(fileCore, zap.AddCaller()), nil
}
