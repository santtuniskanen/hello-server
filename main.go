package main

import (
	"log"
	"os"
	"time"
)

func main() {
	f, err := os.OpenFile("/var/log/hello-go.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	logger := log.New(f, "", log.LstdFlags)
	for {
		logger.Println("Hello, world!")
		time.Sleep(10 * time.Second)
	}
}
