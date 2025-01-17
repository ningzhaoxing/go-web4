package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"userManageSystem-blog/src/pkg/globals"
)

var configPath = "src/configs/dev.yml"

func InitConfig() (*globals.Config, error) {
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("config Load error %s \n", err.Error())
		return nil, err
	}

	var config globals.Config
	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("config bind error %s \n", err.Error())
		return nil, err
	}
	return &config, nil
}
