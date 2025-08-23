package config

import (
	"errors"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
}

type ServerConfig struct {
	Port    string
	RunMode string
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SSLMode  bool
}

type RedisConfig struct {
	Host               string
	Port               string
	User               string
	Password           string
	Db                 string
	MinIdleConnections int
	PoolSize           int
	PoolTimeout        int
}

func GetConfig() *Config {
	cfgPath := getConfigPath(os.Getenv("APP_ENV"))
	v, err := loadConfig(cfgPath, "yml")
	if err != nil {
		log.Fatalf("Unable to load config in Get: %v", err)
	}
	cfg, err := parseConfig(v)
	if err != nil {
		log.Fatalf("Unable to parse config in Get: %v", err)
	}
	return cfg
}

func parseConfig(v *viper.Viper) (*Config, error) {
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Printf("Unable to parse config: %v", err)
		return nil, err
	}
	return &cfg, nil
}

func loadConfig(fileName string, fileType string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType(fileType)
	v.SetConfigName(fileName)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Printf("Unable to load config: %v", err)
			return nil, errors.New("Config file not found")
		}
		return nil, err
	}
	return v, nil
}

func getConfigPath(env string) string {
	switch env {
	case "docker":
		return "config/config-docker"
	case "production":
		return "config/config-production"
	default:
		return "../config/config-development"
	}
}
