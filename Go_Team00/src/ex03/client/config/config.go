package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Host     string         `json:"host"`
	Port     string         `json:"port"`
	Postgres PostgresConfig `json:"postgres"`
	Max      int            `json:"max"`
	K        int            `json:"k"`
}

type PostgresConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DBName   string `json:"dbname"`
}

func DefaultConfigParser() (*Config, error) {
	data, err := os.ReadFile("./config/client.conf")
	if err != nil {
		return nil, err
	}

	conf := &Config{}
	if err := json.Unmarshal(data, conf); err != nil {
		return nil, err
	}

	return conf, nil
}
