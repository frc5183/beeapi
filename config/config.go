package config

import (
	_ "embed"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

//go:embed config.yml
var defaultConfigFile []byte

var Configuration *Config

const (
	DefaultConfigFilePath = "/etc/beeapi/config.yml"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`

	Development bool `yaml:"development"`
}

func GenerateConfig() {
	Configuration = &Config{}

	configFile, err := os.Open(DefaultConfigFilePath)

	if err != nil {
		createConfigFile()
		panic("Created config file. Please edit it and restart the application.")
	}

	err = yaml.NewDecoder(configFile).Decode(Configuration)

	if err != nil {
		panic(err)
	}
}

func createConfigFile() {
	pathParts := strings.Split(DefaultConfigFilePath, "/")
	var path string

	for i := 0; i < len(pathParts)-1; i++ {
		path += pathParts[i] + "/"
	}

	err := os.MkdirAll(path, 0755)
	if err != nil {
		panic(err)
	}

	configFile, err := os.Create(DefaultConfigFilePath)
	if err != nil {
		panic(err)
	}

	_, err = configFile.Write(defaultConfigFile)

	if err != nil {
		panic(err)
	}
}
