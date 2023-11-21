package yaml

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	App    App    `yaml:"app,omitempty" json:"app"`
	Logger Logger `yaml:"logger" json:"logger"`
}

type App struct {
	Name string `yaml:"name,omitempty" json:"name"`
	Port string `yaml:"port,omitempty" json:"port"`
	//ReadTimeOut  int    `yaml:"read_time_out" json:"read_time_out"`
	//WriteTimeOut int    `yaml:"write_time_out" json:"write_time_out"`
}

type Logger struct {
	Enable bool `yaml:"enable" json:"enable"`
}

func NewConfig() (*Config, error) {
	var config *Config

	yfile, err := ioutil.ReadFile("./config/yaml/app.yaml")

	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yfile, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
