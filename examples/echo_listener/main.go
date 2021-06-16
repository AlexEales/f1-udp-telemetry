package main

import (
	"context"
	"f1-udp-telemetry/listener"

	log "github.com/sirupsen/logrus"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	l := listener.NewPacketListener()
	err := l.Start(ctx)
	if err != nil {
		log.WithError(err).Error("error occured while listening for packets")
	}
}
