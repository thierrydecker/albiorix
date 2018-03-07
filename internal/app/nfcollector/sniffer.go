package nfcollector

import (
	"context"
	"fmt"
	"time"

	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket"
)

func Sniffer(snifferContext context.Context, exitSnifferChannel chan string, messageChannel chan LoggingMessage) {
	message := LoggingMessage{}
	message.message = fmt.Sprintf("Sniffer starting...")
	message.logLevel = "INFO"
	messageChannel <- message

	var (
		device            = "enp0s3"
		snapshotLen int32 = 1024
		promiscuous       = true
		err         error
		timeout     = 100 * time.Millisecond
		handle      *pcap.Handle
	)

	handle, err = pcap.OpenLive(device, snapshotLen, promiscuous, timeout)
	if err != nil {
		message := LoggingMessage{}
		message.message = fmt.Sprintf("Sniffer could not start")
		message.logLevel = "FATAL"
		messageChannel <- message
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {
		if packet != nil {
			message = LoggingMessage{}
			message.message = fmt.Sprintf("Captured one packet...")
			message.logLevel = "DEBUG"
			messageChannel <- message
		}
		select {
		case <-snifferContext.Done():
			message = LoggingMessage{}
			message.message = fmt.Sprintf("Sniffer stopped")
			message.logLevel = "INFO"
			messageChannel <- message
			exitSnifferChannel <- "Done"
			return
		default:
		}
	}
}
