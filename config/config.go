package config

import "github.com/spf13/viper"

type Config struct {
	API   *API   `json:"api"`
	MySQL *MySQL `json:"mysql"`
}

type MySQL struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

type API struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func NewConfig() *Config {
	var config Config
	viper.AddConfigPath("./config")
	viper.SetConfigName("config-local")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}
	return &config
}
