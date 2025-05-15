package config

import "os"

type Db struct {
	host string
	port string
	user string
	pass string
	name string
}

type Redis struct {
	host string
	port string
	db   int
}

type Telegram struct {
	token string
}

type Config struct {
	DB Db
	RD Redis
	TG Telegram
}

func Load() *Config {
	return &Config{
		DB: Db{
			host: getEnv("DB_HOST", "localhost"),
			port: getEnv("DB_PORT", "5432"),
			user: getEnv("DB_USER", "aleksey"),
			pass: getEnv("DB_PASSWORD", "password"),
			name: getEnv("DB_NAME", "service"),
		},
		RD: Redis{
			host: getEnv("REDIS_HOST", "localhost"),
			port: getEnv("REDIS_PORT", "6379"),
			db:   0,
		},
		TG: Telegram{
			token: getEnv("TELEGRAM_TOKEN", ""),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
