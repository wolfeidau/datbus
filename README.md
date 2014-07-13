# datbus

This library includes a small simple bus for use in micro services. It currently uses
mqtt as a transport as it is also simple and lightweight.

# Usage

To use this library just import it and configure the bus with an `URL` and a unique
`ClientId` for your service.

```go
url, err := url.Parse("tcp://guest:guest@localhost:1883")

if err != nil {
  logger.Errorf(err)
  os.Exit(2)
}

bus, err = datbus.NewBus(&datbus.Configuration{MqttUrl: url, ClientId: "testapp"})

if err != nil {
  logger.Errorf(err)
  os.Exit(2)
}

err = bus.Connect()

if err != nil {
  logger.Errorf(err)
  os.Exit(2)
}
```

# Licence

Copyright (c) 2014 Mark Wolfe
Licensed under the MIT license.
