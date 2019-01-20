package errors

import (
	"fmt"
	"log"
	"os"

	"github.com/nilorg/sdk/errors"

	"github.com/spf13/viper"
)

var (
	eviper *viper.Viper
)

func init() {
	eviper = viper.New()
	eviper.SetConfigFile("./error.toml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件错误：%s", err.Error())
		os.Exit(-1)
	} else {
		viper.WatchConfig()
	}
}

// GetBusinessError 获取错误信息
func GetBusinessError(code int) *errors.BusinessError {
	msg := eviper.GetString(fmt.Sprintf("business.e%d", code))
	return errors.New(code, msg)
}
