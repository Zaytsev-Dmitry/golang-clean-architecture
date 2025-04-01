package pkg

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type AppConfig struct {
	Server struct {
		Port string `yaml:"port"`
		Name string `yaml:"name"`
	} `yaml:"server"`
	Database struct {
		Host         string `yaml:"host"`
		Username     string `yaml:"userName"`
		Password     string `yaml:"password"`
		DataBaseName string `yaml:"dataBaseName"`
		Dialect      string `yaml:"dialect"`
		Port         string `yaml:"port"`
	} `yaml:"database"`
}

func LoadConfig() *AppConfig {
	configPath := os.Getenv("CONFIG_PATH")
	profile := os.Getenv("APP_PROFILE")

	if configPath == "" && profile == "local" {
		configPath = "configs/" + "local.yaml"
	}

	if configPath == "" {
		log.Fatalf("CONFIG_PATH environment variable not set")
	}

	if profile == "" {
		log.Fatalf("CONFIG_PATH environment variable not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("CONFIG_PATH does not exist: %s", configPath)
	}

	var config AppConfig
	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}
	return &config
}
