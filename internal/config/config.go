package config

import (
	"log"
	"sync"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	DBUser      string `env:"DB_USER"`
	DBPassword  string `env:"DB_PASSWORD"`
	DBName      string `env:"DB_NAME"`
	DBHost      string `env:"DB_HOST"`
	DBPort      int    `env:"DB_PORT" envDefault:"5432"`
	DBSSLMode   string `env:"DB_SSLMODE" envDefault:"disable"`
	DatabaseURL string `env:"DATABASE_URL,required"`

	AppPort int    `env:"APP_PORT" envDefault:"8080"`
	AppEnv  string `env:"APP_ENV" envDefault:"local"`
}

var (
	cfg  *Config
	once sync.Once
)

func Load() *Config {
	once.Do(func() {

		_ = godotenv.Load()

		cfg = &Config{}
		if err := env.Parse(cfg); err != nil {
			log.Fatalf("❌ Config Parse Error: %v", err)
		}
	})

	return cfg
}
