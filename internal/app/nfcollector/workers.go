// Copyright 2018 Thierry DECKER. All rights reserved.
// Use of this source code is governed by a APACHE-2.0
// license that can be found in the LICENSE file.

package nfcollector

import (
	"context"
	"fmt"
	"time"
)

func Worker(ctx context.Context, exitChan chan struct{}, name string, interval time.Duration) {
	i := 1
	for {
		fmt.Printf("%s Waiting (%d) for a signal to stop\n", name, i)
		time.Sleep(time.Second * interval)
		i++
		select {
		case <-ctx.Done():
			fmt.Printf("%s received stop signal, exiting\n", name)
			exitChan <- struct{}{}
			return
		default:
		}
	}
}
