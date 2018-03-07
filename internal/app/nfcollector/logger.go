package nfcollector

import (
	"context"
	log "github.com/sirupsen/logrus"
	"os"
	"fmt"
)

func Logger(loggerContext context.Context, exitLoggerChannel chan string, messageChannel chan LoggingMessage) {
	log.SetFormatter(&log.TextFormatter{ForceColors: true, FullTimestamp: true, DisableSorting: true})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	message := LoggingMessage{}
	message.message = fmt.Sprintf("Logger started")
	message.logLevel = "INFO"
	logMessage(message)
	for {
		select {
		case message := <-messageChannel:
			logMessage(message)
		case <-loggerContext.Done():
			if len(messageChannel) == 0 {
				message := LoggingMessage{}
				message.message = fmt.Sprintf("Logger stopped")
				message.logLevel = "INFO"
				logMessage(message)
				exitLoggerChannel <- "Done"
				return
			}
		default:
		}
	}
}

func logMessage(message LoggingMessage) {
	if message.logLevel == "DEBUG" {
		log.Debugf("%s", message.message)
	}
	if message.logLevel == "INFO" {
		log.Infof("%s", message.message)
	}
	if message.logLevel == "WARNING" {
		log.Warnf("%s", message.message)
	}
	if message.logLevel == "ERROR" {
		log.Errorf("%s", message.message)
	}
	if message.logLevel == "FATAL" {
		log.Fatalf("%s", message.message)
	}
	if message.logLevel == "FATAL" {
		log.Panicf("%s", message.message)
	}
}
