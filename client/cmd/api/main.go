package main

import (
	"github.com/arxdsilva/jpar/client/infrastructure/config"
	"github.com/arxdsilva/jpar/client/reader"
)

func main() {
	config := config.Load()
	go reader.StreamFile(config)
	go reader.SendInfo(config)
}
