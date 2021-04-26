package config

import (
	"os"
	"strconv"

	"github.com/arxdsilva/jpar/client/application"
	"github.com/arxdsilva/jpar/client/domains"
	"github.com/kpango/glg"
)

type Config struct {
	Semaphore   chan domains.Port
	PortService domains.PortService
}

func Load() (c Config) {
	mgr := os.Getenv("MAX_GOROUTINES")
	maxGoroutines, err := strconv.Atoi(mgr)
	if err != nil {
		glg.Fatal("could not load max go routines env var: ", err.Error())
	}
	if maxGoroutines == 0 {
		maxGoroutines = 2
	}
	c.Semaphore = make(chan domains.Port, maxGoroutines)
	c.PortService = application.NewPortService()
	return
}

func Port() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":8888"
}
