package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type PostgresConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Db       string `json:"db"`
}

type Shortener struct {
	AliasLength int `json:"aliasLength`
}

type Config struct {
	PostgresConfig PostgresConfig `json:"postgres"`
	Shortener      Shortener      `json:"shortener`
}

func New(path string) Config {
	content, err := os.Open(path)
	defer content.Close()

	if err != nil {
		log.Fatalf("Could not open %s", path)
	}

	byteValue, _ := io.ReadAll(content)

	var config Config
	json.Unmarshal(byteValue, &config)

	return config
}

func (c Config) GetEnvValue(name string) any {
	return os.Getenv(name)
}
