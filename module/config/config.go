package module

import (
	"log"
	"os"

	"github.com/nilorg/pkg/logger"

	"github.com/spf13/viper"
)

func init() {
	initConfigFile()
	initLog()
}

func initConfigFile() {
	configFile := os.Getenv("CARGOBOAT_CONFIG")
	if configFile == "" {
		configFile = "./cargoboat.toml"
	}
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件错误：%s", err.Error())
	} else {
		viper.WatchConfig()
	}
}
func initLog() {
	// 日志初始化
	logger.Init()
}
