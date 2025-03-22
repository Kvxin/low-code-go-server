package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	MongoURI      string
	RedisAddr     string
	RedisPassword string
	JWTSecret     string
	SMTPHost      string
	SMTPPort      string
	SMTPUser      string
	SMTPPassword  string
}

var cfg Config

func Load() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	cfg = Config{
		MongoURI:      os.Getenv("DATABASE_SERVER_URL"),
		RedisAddr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		JWTSecret:     os.Getenv("JWT_SECRET"),
		SMTPHost:      os.Getenv("EMAIL_HOST"),
		SMTPPort:      os.Getenv("EMAIL_PORT"),
		SMTPUser:      os.Getenv("EMAIL_USER"),
		SMTPPassword:  os.Getenv("EMAIL_PASS"),
	}

	return nil
}

func Get() *Config {
	return &cfg
} 