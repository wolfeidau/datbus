package main

import (
	"net/url"
	"os"
	"os/signal"

	"github.com/juju/loggo"
	datbus "github.com/wolfeidau/datbus"
)

func main() {

	logger := loggo.GetLogger("bus")

	logger.Infof("Started service")

	url, _ := url.Parse("tcp://guest:guest@localhost:2883")

	bus, _ := datbus.NewBus(&datbus.Configuration{MqttUrl: url, ClientId: "testapp"})

	err := bus.Connect()

	if err != nil {
		logger.Errorf("error connecting to bus %s", err)
	}

	logger.Infof("bus %v", bus)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	// Block until a signal is received.
	s := <-c
	logger.Infof("Got signal:", s)
}
