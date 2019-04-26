package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Port        string   `json:"port"`
	ServiceName string   `json:"serviceName"`
	DB          DBConfig `json:"database"`
}

type DBConfig struct {
	User   string `json:"user"`
	Pass   string `json:"pass"`
	DBName string `json:"dbName"`
}

func LoadConfig(file string) (conf *Config, err error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
