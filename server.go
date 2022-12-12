package main

import (
	"github.com/Bexanderthebex/flight-sorter/api"
	"github.com/brpaz/echozap"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func main() {
	e := echo.New()

	zapLogger, _ := zap.NewProduction()

	e.Use(echozap.ZapLogger(zapLogger))
	e.Use(middleware.Recover())

	e.POST("/calculate", api.SortFlights)

	e.Logger.Fatal(e.Start(":8080"))
}
