package main

import (
	"os"
	"os/signal"
	"context"
	"github.com/thierrydecker/albiorix/internal/app/nfcollector"
	"time"
)

func main() {
	exitLoggerChannel := make(chan string)
	exitSnifferChannel := make(chan string)
	messageChannel := make(chan nfcollector.LoggingMessage)
	loggerContext, loggerCancel := context.WithCancel(context.Background())
	snifferContext, snifferCancel := context.WithCancel(context.Background())
	go nfcollector.Logger(loggerContext, exitLoggerChannel, messageChannel)
	go nfcollector.Sniffer(snifferContext, exitSnifferChannel, messageChannel)
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	go func() {
		select {
		case <-signalChannel:
			snifferCancel()
			time.Sleep(time.Millisecond * 1000)
			loggerCancel()
			return
		}
	}()
	<-exitSnifferChannel
	<-exitLoggerChannel
}
