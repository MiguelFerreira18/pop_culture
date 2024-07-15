package config

import (
	"log"
	"time"

	"github.com/joeshaw/envdecode"
)

type DbConfig struct {
	User         string `env:"DATABASE_USER,required"`
	Password     string `env:"DATABASE_PASSWORD,required"`
	UrlAdress    string `env:"URL_ADDRESS,required"`
	UrlPort      int    `env:"URL_PORT,required"`
	DatabaseName string `env:"DATABASE_NAME,required"`
	Locality     string `env:"DATABASE_LOCALITY,required"`
}
type ServerConfig struct {
	Port         string        `env:"SERVER_PORT,required"`
	ReadTimeout  time.Duration `env:"SERVER_READ_TIMEOUT,required"`
	WriteTimeout time.Duration `env:"SERVER_WRITE_TIMEOUT,required"`
	IdleTimeout  time.Duration `env:"SERVER_IDLE_TIMEOUT,required"`
	Debug        bool          `env:"SERVER_DEBUG,required"`
}
type Config struct {
	Database DbConfig
	Server   ServerConfig
}

func New() *Config {
	var config Config
	if err := envdecode.StrictDecode(&config); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}
	return &config
}
