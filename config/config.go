package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		Name string `yaml:"name`
	} `yaml:"database"`

	JWT struct {
		SecretKey string `yaml:"secretKey`
	} `yaml:"JWT"`

	Server struct {
		Port string `yaml:"port`
	} `yaml:"server"`
}

var AppConfig Config

// function to load all the configuration set on the config.yaml file
func LoadConfig(filename string) {
	yamlFile, err := os.ReadFile(filename)

	if err != nil {
		log.Fatalf("Error reading the config file: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &AppConfig)
	if err != nil {
		log.Fatalf("Error in unmarshalling the config %v", err)
	}
}
