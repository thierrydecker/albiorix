// Copyright 2018 Thierry DECKER. All rights reserved.
// Use of this source code is governed by a APACHE-2.0
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/thierrydecker/albiorix/internal/app/nfcollector"
)

func main() {

	c1, cancel := context.WithCancel(context.Background())
	exitCh := make(chan struct{})

	go nfcollector.Worker(c1, exitCh, "worker #1", 1)
	go nfcollector.Worker(c1, exitCh, "worker #2", 2)
	go nfcollector.Worker(c1, exitCh, "worker #3", 3)
	go nfcollector.Worker(c1, exitCh, "worker #4", 4)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)
	go func() {
		select {
		case <-signalCh:
			cancel()
			return
		}
	}()

	for i := 0; i < 4; i++ {
		<-exitCh
	}
}
