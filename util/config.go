package util

import (
	"time"

	"github.com/spf13/viper"
)

//config loading through viper.
type Config struct{
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	TokenAssymetricKey string `mapstructure:"TOKEN_ASSYMETRIC_KEY"`
	AcccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRY_DURATION"`
}

func LoadConfig(path string)(config Config, err error){
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil{
		return
	}

	err = viper.Unmarshal(&config)
	return
}