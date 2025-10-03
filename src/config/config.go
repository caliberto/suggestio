package config

import "os"

type Config struct {
	BotToken string
	Port     string
}

func Load() *Config {
	return &Config{
		BotToken: os.Getenv("BOT_TOKEN"),
		Port:     getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return fallback
}
