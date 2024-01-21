package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Postgres PostgresConfig `json:"postgres"`
}

type PostgresConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"hostname"`
	Port     string `json:"port"`
	DBName   string `json:"dbname"`
}

func ConfigParser() (*Config, error) {
	data, err := os.ReadFile("client.conf")
	if err != nil {
		return nil, err
	}

	conf := &Config{}
	if err := json.Unmarshal(data, conf); err != nil {
		return nil, err
	}

	return conf, nil
}
