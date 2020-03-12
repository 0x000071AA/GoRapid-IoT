package mqtt

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type MQTTClient struct {
	c        MQTT.Client
	quos     byte
	retained bool
}

func MQTTInit(broker, user, password string) (MQTTClient, error) {
	opts := MQTTConfig(broker, user, password)
	clientPub := MQTTClient{
		c:        MQTT.NewClient(opts),
		quos:     0,
		retained: false,
	}

	if tokenPub := clientPub.c.Connect(); tokenPub.Wait() && tokenPub.Error() != nil {
		return clientPub, tokenPub.Error()
	}

	return clientPub, nil
}

func (mqtt MQTTClient) Publish(topic string, payload interface{}) {
	tokenPub := mqtt.c.Publish(topic, mqtt.quos, mqtt.retained, payload)
	tokenPub.Wait()
}

func (mqtt MQTTClient) Subscribe(topic string, cb MQTT.MessageHandler) error {
	tokenPub := mqtt.c.Subscribe(topic, mqtt.quos, cb)
	if tokenPub.Wait() && tokenPub.Error() != nil {
		return tokenPub.Error()
	}
	return nil
}

func (mqtt MQTTClient) Unsubscribe(topics ...string) error {
	if t := mqtt.c.Unsubscribe(topics...); t.Wait() && t.Error() != nil {
		return t.Error()
	}
	return nil
}
