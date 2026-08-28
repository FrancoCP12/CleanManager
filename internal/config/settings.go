package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file: %v", err)
		return nil, err
	}

	config := &Config{
		DBUser:     viper.GetString("database.user"),
		DBPassword: viper.GetString("database.password"),
		DBHost:     viper.GetString("database.host"),
		DBPort:     viper.GetString("database.port"),
		DBName:     viper.GetString("database.name"),
	}

	return config, nil
}
