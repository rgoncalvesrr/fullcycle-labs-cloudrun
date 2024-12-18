package configs

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	WeatherApiKey string `mapstructure:"WEATHER_API_KEY"`
	WeatherApiUrl string `mapstructure:"WEATHER_API_URL"`
	CepApiUrl     string `mapstructure:"CEP_API_URL"`
}

func NewConfig() *Config {
	//cfgFile := filepath.Join(path, ".env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AddConfigPath("../..")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	e := viper.ReadInConfig()
	if e != nil {
		log.Fatal("Can't find the file .env : ", e)
	}

	var result Config

	e = viper.Unmarshal(&result)
	if e != nil {
		log.Fatal("Can't unmarshal the file .env : ", e)
	}

	return &result
}
