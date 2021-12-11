package config

import "github.com/spf13/viper"

type Config struct {
	App struct {
		Name string `json:"name"`
		Host string `json:"host"`
	} `json:"app"`
	Database struct {
		Mysql string `json:"mysql"`
	} `json:"database"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
