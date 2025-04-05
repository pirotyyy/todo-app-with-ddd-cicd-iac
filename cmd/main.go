package main

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pirotyyy/todo-app-with-ddd-cicd-iac/internal/config"
	"github.com/pirotyyy/todo-app-with-ddd-cicd-iac/internal/logger"
)

func main() {
	ctx := context.Background() // TODO: ちゃんとやる

	a := "unused"

	conf, err := config.Load()
	if err != nil {
		logger.Error(ctx, err.Error())
	}
	slog.SetDefault(slog.New(logger.NewHandler(conf.Log)))

	e := echo.New()

	e.Use(middleware.Recover())

	// health check
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
