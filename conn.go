package datbus

import (
	"fmt"
	"net/url"

	"github.com/juju/loggo"

	mqtt "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
)

type Message struct {
	mqtt.Message
}

// Wraps the mqtt connection and deals with reconnects, recovery and generally
// hides this mess
type NinjaConnection struct {
	mqtt *mqtt.MqttClient
	log  loggo.Logger
}

func Connect(url *url.URL, clientId string) (*NinjaConnection, error) {

	logger := loggo.GetLogger(fmt.Sprintf("%s.%s", "conn", clientId))

	opts :=
		mqtt.NewClientOptions().SetBroker(url.String()).SetClientId(clientId).SetTraceLevel(mqtt.Off)

	conn := &NinjaConnection{mqtt: mqtt.NewClient(opts), log: logger}

	_, err := conn.mqtt.Start()

	if err != nil {
		logger.Errorf("Failed to connect to mqtt server %s - %s", url.String(), err)

		// TODO reconnect required
		return nil, err

	} else {
		logger.Infof("Connected to %s", url.String())
	}

	return conn, nil
}

type Handler func(msg Message, conn *NinjaConnection)

func (conn *NinjaConnection) SubscribeFunc(topicPattern string, handler Handler) error {

	topicFilter, err := mqtt.NewTopicFilter(topicPattern, 0)

	if err != nil {
		return err
	}

	receipt, err := conn.mqtt.StartSubscription(func(client *mqtt.MqttClient, msg mqtt.Message) {

		handler(Message{msg}, conn)
	}, topicFilter)

	<-receipt
	conn.log.Infof("subscribed to %s", topicPattern)

	return err

}
