package config

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"log"
	"os"
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
type KeyPaths struct {
	PublicKeyPath  string `env:"PUBLIC_KEY_PATH,required"`
	PrivateKeyPath string `env:"PRIVATE_KEY_PATH,required"`
}

type RSAKeys struct {
	PrivateKey rsa.PrivateKey
	PublicKey  rsa.PublicKey
}

type Config struct {
	Database DbConfig
	Server   ServerConfig
	KeyPaths KeyPaths
	Keys     RSAKeys
}

func New() *Config {
	var config Config
	if err := envdecode.StrictDecode(&config); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}

	return &config
}

func LoadPrivateKey(path string) (*rsa.PrivateKey, error) {
	keyData, err := os.ReadFile(path)
	if err != nil {
		return nil, err

	}

	block, _ := pem.Decode(keyData)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("failed to decode PEM block containing private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func LoadPublicKey(path string) (*rsa.PublicKey, error) {

	keyData, err := os.ReadFile(path)
	if err != nil {
		return nil, err

	}

	block, _ := pem.Decode(keyData)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, errors.New("failed to decode PEM block containing private key")
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return publicKey.(*rsa.PublicKey), nil
}
