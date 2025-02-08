package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		Server     ServerConfig
		Database   DatabaseConfig
		Redis      RedisConfig
		ProID      ProIDService
		ApiKey     string
		ProIDOther map[string]*ProIDOther
	}
	ProIDOther struct {
		ClientID    string
		SecretKey   string
		RedirectURL string
	}
	ProIDService struct {
		Endpoint  string
		BaseURL   string
		GrantType string
	}
	ServerConfig struct {
		Host string
		Port string // Server port
	}

	DatabaseConfig struct {
		Host     string // Database host
		Port     string // Database port
		User     string // Database user
		Password string // Database password
		DBName   string // Database name
	}

	RedisConfig struct {
		Host string // Redis host
		Port string // Redis port
	}
)

// Load loads configuration from .env file
func (c *Config) Load() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}
	c.ProIDOther = make(map[string]*ProIDOther, 3)
	c.ProIDOther["web"] = &ProIDOther{}
	c.ProIDOther["admin"] = &ProIDOther{}
	c.ProIDOther["mobile"] = &ProIDOther{}

	// Fetch and validate required environment variables
	requiredVars := map[string]*string{
		//gRPC server
		"SERVER_HOST": &c.Server.Host,
		"SERVER_PORT": &c.Server.Port,
		//database
		"DB_PORT":     &c.Database.Port,
		"DB_USER":     &c.Database.User,
		"DB_PASSWORD": &c.Database.Password,
		"DB_NAME":     &c.Database.DBName,
		"DB_HOST":     &c.Database.Host,
		"REDIS_HOST":  &c.Redis.Host,
		"REDIS_PORT":  &c.Redis.Port,
		//proid configs
		"PROID_GRANT_TYPE": &c.ProID.GrantType,
		"PROID_BASE_URL":   &c.ProID.BaseURL,
		"PROID_ENDPOINT":   &c.ProID.Endpoint,

		"PROID_CLIENT_ID":     &c.ProIDOther["web"].ClientID,
		"PROID_CLIENT_SECRET": &c.ProIDOther["web"].SecretKey,
		"PROID_REDIRECT_URI":  &c.ProIDOther["web"].RedirectURL,

		"PROID_ADMIN_CLIENT_ID":     &c.ProIDOther["admin"].ClientID,
		"PROID_ADMIN_CLIENT_SECRET": &c.ProIDOther["admin"].SecretKey,
		"PROID_ADMIN_REDIRECT_URI":  &c.ProIDOther["admin"].RedirectURL,

		"API_KEY": &c.ApiKey,
	}

	for envVar, fieldPtr := range requiredVars {
		value := os.Getenv(envVar)
		if value == "" {
			return fmt.Errorf("missing required environment variable: %s", envVar)
		}
		*fieldPtr = value
	}
	// Add colon prefix to the server port

	return nil
}

// New creates a new Config instance and loads the configuration
func New() (*Config, error) {
	var config Config
	if err := config.Load(); err != nil {
		return nil, err
	}
	return &config, nil
}
