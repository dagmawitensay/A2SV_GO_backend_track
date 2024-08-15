package config

import (
	"fmt"
	"os"
	domain "task_manager_api_clean_architecture/Domain"

	"github.com/spf13/viper"
)




func LoadConfig(path string) (config domain.Config, err error) {
  viper.AddConfigPath(path)
  viper.SetConfigType("env")
  viper.SetConfigName(".env")

  viper.AutomaticEnv()
  p, err := os.Getwd()
  fmt.Println("printing workin directory", p)

  err = viper.ReadInConfig()
  if err != nil {
    return
  }

  err = viper.Unmarshal(&config)
  return
}

	 