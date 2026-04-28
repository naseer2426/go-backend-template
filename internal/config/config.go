package config

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

// ServerConfig holds required HTTP server settings.
type ServerConfig struct {
	Port string `validate:"required"`
}

// DatabaseConfig holds optional database settings — zero URL means no database is used.
type DatabaseConfig struct {
	URL string `validate:"omitempty"`
}

// Config is the root application configuration.
type Config struct {
	Server   ServerConfig   `validate:"required"`
	Database DatabaseConfig
}

var (
	instance *Config
	once     sync.Once
)

// MustLoad loads .env if present, reads environment into a [Config], validates it, and stores a singleton.
// It panics if .env exists but cannot be read, or if validation fails (e.g. missing required fields).
func MustLoad() *Config {
	once.Do(func() {
		if err := loadDotEnv(); err != nil {
			panic(err)
		}
		cfg := fromEnv()
		applyDefaults(cfg)
		if err := validateConfig(cfg); err != nil {
			panic(fmt.Errorf("config validation failed: %w", err))
		}
		instance = cfg
	})
	return instance
}

// Get returns the loaded config singleton. Panics if [MustLoad] was not called first.
func Get() *Config {
	if instance == nil {
		panic("config not loaded: call config.MustLoad first")
	}
	return instance
}

func loadDotEnv() error {
	if err := godotenv.Load(); err != nil {
		if !os.IsNotExist(err) {
			log.Printf("warning: could not load .env: %v", err)
			return err
		}
	}
	return nil
}

func fromEnv() *Config {
	return &Config{
		Server: ServerConfig{
			Port: strings.TrimSpace(os.Getenv("PORT")),
		},
		Database: DatabaseConfig{
			URL: strings.TrimSpace(os.Getenv("DATABASE_URL")),
		},
	}
}

func applyDefaults(cfg *Config) {
	if cfg.Server.Port == "" {
		cfg.Server.Port = "8080"
	}
}

func validateConfig(cfg *Config) error {
	v := validator.New()
	return v.Struct(cfg)
}
