package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		Server   ServerConfig
		Database DatabaseConfig
		Redis    RedisConfig
		Apis     APIs
	}

	ServerConfig struct {
		Port string
	}

	DatabaseConfig struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
	}

	RedisConfig struct {
		Host string
		Port string
	}

	APIs struct {
		WeatherApi  string
		CurrencyApi string
	}
)

func (c *Config) Load() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	c.Server.Port = ":" + os.Getenv("SERVER_PORT")
	c.Database.Host = os.Getenv("DB_HOST")
	c.Database.Port = os.Getenv("DB_PORT")
	c.Database.User = os.Getenv("DB_USER")
	c.Database.Password = os.Getenv("DB_PASSWORD")
	c.Database.DBName = os.Getenv("DB_NAME")
	c.Redis.Host = os.Getenv("REDIS_HOST")
	c.Redis.Port = os.Getenv("REDIS_PORT")
	c.Apis.CurrencyApi = os.Getenv("CURRENCY_API")
	c.Apis.WeatherApi = os.Getenv("WEATHER_API")

	return nil
}

func New() (*Config, error) {
	var config Config
	if err := config.Load(); err != nil {
		return nil, err
	}
	return &config, nil
}

// REDIS_URI=redis_uri
