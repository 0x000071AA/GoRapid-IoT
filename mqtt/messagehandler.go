package mqtt

import (
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func DefaultMessageHandler(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", msg.Topic(), msg.Payload())
}
