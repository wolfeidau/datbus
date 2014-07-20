package datbus

import (
	"git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
	"github.com/juju/loggo"
)

var logger = loggo.GetLogger("bus")

type Bus struct {
	config *Configuration
	conn   *BusConnection
}

func NewBus(config *Configuration) (*Bus, error) {
	return &Bus{config: config}, nil
}

// connect to the mqtt broker configured when the bus was created.
func (bus *Bus) Connect() error {

	conn, err :=
		Connect(bus.config.MqttUrl, bus.config.ClientId)

	if err != nil {
		return err
	}
	bus.conn = conn

	return nil
}

type Handler func(msg Message, conn *BusConnection)

func (bus *Bus) SubscribeFunc(topicPattern string, handler Handler) error {

	topicFilter, err := mqtt.NewTopicFilter(topicPattern, 0)

	if err != nil {
		return err
	}

	receipt, err := bus.conn.mqtt.StartSubscription(func(client *mqtt.MqttClient, msg mqtt.Message) {

		handler(Message{msg}, bus.conn)
	}, topicFilter)

	<-receipt
	bus.conn.log.Infof("subscribed to %s", topicPattern)

	return err

}
