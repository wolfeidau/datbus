package datbus

import "net/url"

// holds all standard configuration for bus
type Configuration struct {
	MqttUrl  *url.URL
	ClientId string
}
