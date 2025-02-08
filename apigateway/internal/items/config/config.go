/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-09-21 19:53:48
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-12-26 01:29:00
 * @FilePath: /sfere_backend/internal/items/config/config.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		Server         ServerConfig
		Database       DatabaseConfig
		Redis          RedisConfig
		Apis           APIs
		Search         Search
		SearchServices gRpcService
		AdminServices  gRpcService
		AuthService    gRpcService
		FrontEnd       FrontEnd
		FrontEndAdmin  FrontEndAdmin
		ApiKey         string
	}
	FrontEndAdmin struct {
		RedirectURL string
	}
	FrontEnd struct {
		RedirectURL string
	}
	gRpcService struct {
		Name string
		Port string
		Host string
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

	c.Server.Port = os.Getenv("SERVER_HOST") + ":" + os.Getenv("SERVER_PORT")

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
	c.SearchServices.Port = os.Getenv("gRPC_SEARCH_SERVER_PORT")
	c.SearchServices.Host = os.Getenv("gRPC_SEARCH_SERVER_HOST")

	c.AuthService.Host = os.Getenv("gRPC_AUTH_SERVER_HOST")
	c.AuthService.Port = os.Getenv("gRPC_AUTH_SERVER_PORT")

	c.AdminServices.Port = os.Getenv("gRPC_ADMIN_SERVER_PORT")
	c.AdminServices.Host = os.Getenv("gRPC_ADMIN_SERVER_HOST")

	c.FrontEnd.RedirectURL = os.Getenv("FRONT_REDIRECT_URL")
	c.FrontEndAdmin.RedirectURL = os.Getenv("FRONT_ADMIN_REDIRECT_URL")

	c.ApiKey = os.Getenv("API_KEY")
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
