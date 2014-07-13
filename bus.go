package datbus

import (
	"github.com/juju/loggo"
)

type Bus struct {
	config *Configuration
	conn   *NinjaConnection
	log    loggo.Logger
}

func NewBus(config *Configuration) (*Bus, error) {
	return &Bus{config: config, log: loggo.GetLogger("bus")}, nil
}

func (bus *Bus) Connect() error {

	conn, err :=
		Connect(bus.config.MqttUrl, bus.config.ClientId)

	if err != nil {
		return err
	}
	bus.conn = conn

	return nil
}
