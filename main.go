package main

import (
	"github.com/onedss/RTSPtoWSMP4f/core"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	go core.ServeHTTP()
	go core.ServeStreams()
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		log.Println(sig)
		done <- true
	}()
	log.Println("Server Start Awaiting Signal")
	<-done
	log.Println("Exiting")
}
