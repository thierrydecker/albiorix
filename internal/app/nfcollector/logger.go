package nfcollector

import (
	"context"
	log "github.com/sirupsen/logrus"
	"os"
)

func Logger(loggerContext context.Context, exitLoggerChannel chan string, messageChannel chan string) {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.Debugf("Logging worker started")
	for {
		select {
		case message := <-messageChannel:
			log.Debugf("%s", message)
		case <-loggerContext.Done():
			if len(messageChannel) == 0 {
				log.Debugf("Logger worker stopped")
				exitLoggerChannel <- "Done"
				return
			}
		default:
		}
	}
}
