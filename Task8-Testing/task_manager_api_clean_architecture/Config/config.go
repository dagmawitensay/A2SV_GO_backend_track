package config

import (
	domain "task_manager_api_clean_architecture/Domain"

	"github.com/spf13/viper"
)




func LoadConfig(path string) (config domain.Config, err error) {
  viper.AddConfigPath(path)
  viper.SetConfigType("env")
  viper.SetConfigName(".env")

  viper.AutomaticEnv()

  err = viper.ReadInConfig()
  if err != nil {
    return
  }

  err = viper.Unmarshal(&config)
  return
}

	 