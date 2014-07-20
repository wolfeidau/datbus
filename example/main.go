package main

import (
	"net/url"
	"os"
	"os/signal"

	"github.com/juju/loggo"
	bus "github.com/wolfeidau/datbus"
)

var logger = loggo.GetLogger("example")

func SysHandler(msg bus.Message, conn *bus.BusConnection) {
	logger.Infof("received message %s %s", msg.Topic(), msg.Payload())
}

func main() {
	// set the default level
	loggo.GetLogger("").SetLogLevel(loggo.TRACE)

	logger.Infof("Started service")

	url, _ := url.Parse("tcp://guest:guest@localhost:1883")

	bus, _ := bus.NewBus(&bus.Configuration{MqttUrl: url, ClientId: "testapp"})

	err := bus.Connect()

	if err != nil {
		logger.Errorf("error connecting to bus %s", err)
	}

	bus.SubscribeFunc("$SYS/#", SysHandler)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	// Block until a signal is received.
	s := <-c
	logger.Infof("Got signal: %v", s)
}
