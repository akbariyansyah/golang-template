package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func RegisterServer(lc fx.Lifecycle, app *echo.Echo, port int) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := app.Start(fmt.Sprintf(":%d", port)); err != nil {
				log.Fatalf("failed to start server: %v", err)
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown(ctx)
		},
	})
}

var ServerModule = fx.Options(
	fx.Invoke(RegisterServer),
)
