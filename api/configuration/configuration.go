package configuration

import (
	"fmt"

	"github.com/spf13/viper"
)

type DbConfig struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUsername string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASS"`
	DBPort     int16  `mapstructure:"DB_PORT"`
}

func (config DbConfig) GetDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", config.DBHost, config.DBUsername, config.DBPassword, config.DBName, config.DBPort)
}

func GetConfig() DbConfig {
	var config DbConfig
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Reading configuration failed: %w", err))
	}
	err := viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("Configuration deserialization failure: %w", err))
	}
	return config
}
