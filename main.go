package main

import (
	"gold/src/logger"
	"gold/src/server"
)

func main() {
	log := logger.Get()
	defer func() {
		_ = log.Sync()
	}()

	if err := server.Start(); err != nil {
		panic(err)
	}
}
