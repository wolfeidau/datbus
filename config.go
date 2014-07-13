package datbus

import "net/url"

// holds all standard configuration for bus and other pieces which is common
// to all ninja stuff

type Configuration struct {
	MqttUrl  *url.URL
	ClientId string
}
