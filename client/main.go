package main

import (
	"os"
	"strconv"
	"sync"

	"github.com/arxdsilva/jpar/client/reader"
	"github.com/kpango/glg"
)

var maxGoroutines int

type Config struct {
}

func main() {
	load()
	semaphore := make(chan reader.CityPort, maxGoroutines)
	var wg sync.WaitGroup
	wg.Add(2)
	go reader.StreamFile(semaphore)
	go reader.SendInfo(semaphore)
	wg.Wait()
}

func load() {
	mgr := os.Getenv("MAX_GOROUTINES")
	var err error
	maxGoroutines, err = strconv.Atoi(mgr)
	if err != nil {
		glg.Fatal("could not load max go routines env var: ", err.Error())
	}
	if maxGoroutines == 0 {
		maxGoroutines = 2
	}
}
