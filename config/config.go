package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var C config

func init() {
	var err error
	configBytes, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic("failed to read config file:\n" + err.Error())
	}
	err = yaml.Unmarshal(configBytes, &C)
	if err != nil {
		panic("failed to unmarshal config bytes:\n" + err.Error())
	}
}

type config struct {
	HttpListenAddress string `yaml:"http_listen_address"`
}
