package config

import (
	"os"
	"strconv"

	"github.com/arxdsilva/jpar/client/domains"
	"github.com/kpango/glg"
)

type Config struct {
	Semaphore chan domains.Port
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
	return
}
