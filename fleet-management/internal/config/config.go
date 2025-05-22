package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port        string `mapstructure:"PORT"`
	DatabaseURL string `mapstructure:"DATABASE_URL"`
	JWTSecret   string `mapstructure:"JWT_SECRET"`
}

func LoadConfig() (Config, error) {
	var config Config
	viper.SetDefault("PORT", "8001")
	viper.SetDefault("DATABASE_URL", "postgres://postgres:postgres@chetapp-fleet-management-db:5432/chetapp-fleet-management-db")
	viper.SetDefault("JWT_SECRET", "your_jwt_secret")

	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}
