/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-11-22 23:02:16
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-12-21 03:11:23
 * @FilePath: /admin/internal/items/config/config.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		Server   ServerConfig
		Database DatabaseConfig
		Redis    RedisConfig
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

	// Fetch and validate required environment variables
	requiredVars := map[string]*string{
		"SERVER_HOST": &c.Server.Host,
		"SERVER_PORT": &c.Server.Port,
		"DB_PORT":     &c.Database.Port,
		"DB_USER":     &c.Database.User,
		"DB_PASSWORD": &c.Database.Password,
		"DB_NAME":     &c.Database.DBName,
		"DB_HOST":     &c.Database.Host,
		"REDIS_HOST":  &c.Redis.Host,
		"REDIS_PORT":  &c.Redis.Port,
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
