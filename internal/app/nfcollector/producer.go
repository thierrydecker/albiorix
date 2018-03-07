package nfcollector

import (
	"context"
	"fmt"
)

func Producer(producerContext context.Context, exitProducerChannel chan string, messageChannel chan string) {
	messageChannel <- fmt.Sprintf("Producer started")
	for i := 0; ; i++ {
		select {
		case <-producerContext.Done():
			messageChannel <- fmt.Sprintf("Produced %d items", i)
			messageChannel <- fmt.Sprintf("Producer stopped")
			exitProducerChannel <- "Done"
			return
		default:
			messageChannel <- fmt.Sprintf("Produced message %d", i)
		}
	}
}
