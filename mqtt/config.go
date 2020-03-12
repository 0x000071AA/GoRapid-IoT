package mqtt

import (
	"crypto/sha1"
	"fmt"
	"math/rand"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var characterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = characterRunes[rand.Intn(len(characterRunes))]
	}
	return string(b)
}

func NewSHA1Hash(n int) string {
	randString := RandomString(n)

	hash := sha1.New()
	hash.Write([]byte(randString))
	bs := hash.Sum(nil)

	return fmt.Sprintf("%x", bs)
}

func MQTTConfig(broker, username, password string) *MQTT.ClientOptions {
	fmt.Println("pubsub : Pub/Sub initializing...")

	optsPub := MQTT.NewClientOptions()
	optsPub.AddBroker(broker)
	optsPub.SetUsername(username)
	optsPub.SetPassword(password)
	optsPub.SetClientID(fmt.Sprintf("client-%s", NewSHA1Hash(8)))
	optsPub.SetCleanSession(false)
	optsPub.SetDefaultPublishHandler(DefaultMessageHandler)

	return optsPub
}
