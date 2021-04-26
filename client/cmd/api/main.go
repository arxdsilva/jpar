package main

import (
	"github.com/arxdsilva/jpar/client/infrastructure/config"
	"github.com/arxdsilva/jpar/client/interfaces"
)

func main() {
	config := config.Load()
	interfaces.Run(config)
}
