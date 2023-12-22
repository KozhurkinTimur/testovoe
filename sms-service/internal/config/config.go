package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env string `yaml:"env" env-defualt:"local"`
	//StoragePath string `yaml:"storage_path" env-required:"true"`
}

// type HTTPServer struct {
// 	Address     string        `yaml:"address" env-defualt:"localhost:8082"`
// 	Timeout     time.Duration `yaml:"timeout" env-defualt:"4s"`
// 	IdleTimeout time.Duration `yaml:"idle_timeout" env-defualt:"60s"`
// }

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
