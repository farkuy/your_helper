package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	Port         string        `yaml: "port"`
	Host         string        `yaml: "host"`
	ReadTimeout  time.Duration `yaml: "readTimeout"`
	WriteTimeout time.Duration `yaml: "writeTimeout"`
	Environment  string        `yaml: "environment"`
}

func Init() (*ServerConfig, error) {
	err := godotenv.Load("../../.env")
	if err != nil {
		return nil, fmt.Errorf("ошибка пути нахождения конфигурации: %v", err)
	}

	configPath := os.Getenv("YAML_PATH")
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении конфигурации: %v", err)
	}

	var config ServerConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("ошибка парсинга конфигурации: %v", err)
	}

	return &config, nil
}
