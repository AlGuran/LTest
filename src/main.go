package main

import (
	"LTest/src/http"
	"LTest/src/matchmaker"
	"github.com/mgutz/logxi/v1"
	"os"
	"os/signal"
	"syscall"
)

var (
	signalChan = make(chan os.Signal, 1)
	stopChan   = make(chan struct{})
	systemPort = 1234
	logger     = log.New("config")
)

func main() {
	logger.Info("APP: starting")

	signal.Notify(signalChan, os.Interrupt)
	signal.Notify(signalChan, syscall.SIGTERM)
	go func() {
		<-signalChan
		close(stopChan)
	}()

	go matchmaker.MatchMake()
	http.StartHttp(systemPort)

	<-stopChan
	logger.Info("APP: SHUTDOWN NOW")
	os.Exit(0)
}
