package main

import (
	"github.com/arxdsilva/jpar/client/infrastructure/config"
	"github.com/arxdsilva/jpar/client/reader"
	"github.com/labstack/echo/v4"
)

func main() {
	config := config.Load()
	go reader.StreamFile(config)
	go reader.SendInfo(config)
	e := echo.New()
	// middlewares.Load(e)
	// middlewares.RegisterRoutes(e)
	_ = e
}
