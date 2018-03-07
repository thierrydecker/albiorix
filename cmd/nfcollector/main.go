// Copyright 2018 Thierry DECKER. All rights reserved.
// Use of this source code is governed by a APACHE-2.0
// license that can be found in the LICENSE file.

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
	exitProducerChannel := make(chan string)
	messageChannel := make(chan string)
	loggerContext, loggerCancel := context.WithCancel(context.Background())
	producerContext, producerCancel := context.WithCancel(context.Background())
	go nfcollector.Logger(loggerContext, exitLoggerChannel, messageChannel)
	go nfcollector.Producer(producerContext, exitProducerChannel, messageChannel)
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	go func() {
		select {
		case <-signalChannel:
			producerCancel()
			time.Sleep(time.Millisecond * 500)
			loggerCancel()
			return
		}
	}()
	<-exitProducerChannel
	<-exitLoggerChannel
}
