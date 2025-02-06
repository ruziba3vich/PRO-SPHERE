/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-11-15 19:54:30
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-11-20 22:10:25
 * @FilePath: /searching/internal/items/config/config.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
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
		Search   Search
	}

	ServerConfig struct {
		Port string
		Host string
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

	Search struct {
		Link string
		Key  string
		CxID string
	}
)

func (c *Config) Load() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	c.Server.Port = os.Getenv("SERVER_PORT")
	c.Server.Host = os.Getenv("SERVER_HOST")
	c.Database.Host = os.Getenv("DB_HOST")
	c.Database.Port = os.Getenv("DB_PORT")
	c.Database.User = os.Getenv("DB_USER")
	c.Database.Password = os.Getenv("DB_PASSWORD")
	c.Database.DBName = os.Getenv("DB_NAME")
	c.Redis.Host = os.Getenv("REDIS_HOST")
	c.Redis.Port = os.Getenv("REDIS_PORT")
	c.Apis.CurrencyApi = os.Getenv("CURRENCY_API")
	c.Apis.WeatherApi = os.Getenv("WEATHER_API")
	c.Search.Link = os.Getenv("SEARCH_API")
	c.Search.Key = os.Getenv("API_KEY")
	c.Search.CxID = os.Getenv("CX_ID")

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

func (c *Config) CustomSearch() string {
	return c.Search.Link + "customsearch/v1?" + "cx=" + c.Search.CxID + "&key=" + c.Search.Key
}

func (c *Config) YoutubeSearch() string {
	return c.Search.Link + "youtube/v3/search?" + "cx=" + c.Search.CxID + "&key=" + c.Search.Key
}

func (c *Config) ImagesSearch() string {
	return c.Search.Link + "customsearch/v1?" + "cx=" + c.Search.CxID + "&key=" + c.Search.Key + "&searchType=image"
}
