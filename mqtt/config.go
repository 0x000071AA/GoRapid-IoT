package mqtt

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"os"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"gopkg.in/yaml.v2"
)

type Mqtt struct {
	Client struct {
		ConnectTimeout int    `yaml:"connectTimeout"`
		Username       string `yaml:"username"`
		Password       string `yaml:"password"`
	} `yaml:"client"`
	Broker struct {
		BrokerURL string `yaml:"url"`
	} `yaml:"broker"`
}

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

	seconds := time.Duration(40) * time.Second

	optsPub := MQTT.NewClientOptions()
	optsPub.SetConnectTimeout(seconds)
	optsPub.AddBroker(broker)
	optsPub.SetUsername(username)
	optsPub.SetPassword(password)
	optsPub.SetClientID(fmt.Sprintf("client-%s", NewSHA1Hash(8)))
	optsPub.SetCleanSession(false)
	optsPub.SetDefaultPublishHandler(DefaultMessageHandler)

	return optsPub
}

func MQTTConfigFromYAML(file string) (*MQTT.ClientOptions, error) {
	config := &Mqtt{}

	f, err := os.Open(file)

	defer f.Close()

	if err != nil {
		return nil, err
	}

	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	seconds := time.Duration(config.Client.ConnectTimeout) * time.Second

	optsPub := MQTT.NewClientOptions()
	optsPub.SetConnectTimeout(seconds)
	optsPub.AddBroker(config.Broker.BrokerURL)
	if config.Client.Username != "" && config.Client.Password != "" {
		optsPub.SetUsername(config.Client.Username)
		optsPub.SetPassword(config.Client.Password)
	}
	optsPub.SetClientID(fmt.Sprintf("client-%s", NewSHA1Hash(8)))
	optsPub.SetCleanSession(false)
	optsPub.SetDefaultPublishHandler(DefaultMessageHandler)

	return optsPub, nil
}
