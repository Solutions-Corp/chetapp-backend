package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port                string `mapstructure:"PORT"`
	DefaultUserEmail    string `mapstructure:"DEFAULT_USER_EMAIL"`
	DefaultUserPassword string `mapstructure:"DEFAULT_USER_PASSWORD"`
	DatabaseURL         string `mapstructure:"DATABASE_URL"`
	JWTSecret           string `mapstructure:"JWT_SECRET"`
}

func LoadConfig() (Config, error) {
	var config Config
	viper.SetDefault("PORT", "8000")
	viper.SetDefault("DEFAULT_USER_EMAIL", "admin@chetapp.com")
	viper.SetDefault("DEFAULT_USER_PASSWORD", "qC0sTVXwfJkUg4hDp8yX")
	viper.SetDefault("DATABASE_URL", "postgres://postgres:postgres@auth-db:5432/chetapp-auth-db")
	viper.SetDefault("JWT_SECRET", "secret")

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
