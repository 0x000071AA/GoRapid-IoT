package opcua

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type OpcuaDescription struct {
	Endpoint struct {
		URL      string `yaml:"url"`
		Port     int    `yaml:"port"`
		NodeID   string `yaml:nodeId`
		Policy   string `yaml:"securityPolicy"`
		Mode     string `yaml:"securityMode"`
		Cert     string `yaml:"certificate"`
		Key      string `yaml:"privateKey"`
		Interval int    `yaml:"subscriptionInterval"`
	}
}

type Description struct {
	Endpoint string
	NodeID   string
	Policy   string
	Mode     string
	Interval time.Duration
}

func GetConfigFromYAML(file string) (*Description, error) {
	config := &OpcuaDescription{}
	description := &Description{}

	f, err := os.Open(file)

	defer f.Close()

	if err != nil {
		return nil, err
	}

	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	// check for empty values and set default values
	description.Endpoint = fmt.Sprintf("%v:%d", config.Endpoint.URL, config.Endpoint.Port)

	if config.Endpoint.Cert == "" && config.Endpoint.Key == "" {
		description.Mode = "None"
		description.Policy = "None"
	}

	if config.Endpoint.Interval == 0 {
		description.Interval = time.Duration(40) * time.Second
	} else {
		description.Interval = time.Duration(config.Endpoint.Interval) * time.Second
	}

	return description, nil
}
