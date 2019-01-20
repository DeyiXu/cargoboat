package module

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nilorg/pkg/logger"
	"github.com/spf13/viper"
)

func init() {
	initConfigFile()
	initLog()
	initGin()
}

func initConfigFile() {
	configFile := os.Getenv("CARGOBOAT_CONFIG")
	if configFile == "" {
		configFile = "./cargoboat.toml"
	}
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件错误：%s", err.Error())
		os.Exit(-1)
	}
}

func initLog() {
	// 日志初始化
	logger.Init()
}

func initGin() {
	if viper.GetString("system.mode") == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}
