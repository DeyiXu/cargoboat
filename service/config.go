package service

import (
	"github.com/cargoboat/cargoboat/model"
	"github.com/nilorg/pkg/logger"
)

var (
	configChannel = make(chan model.Config)
	// PubConfigChannel 发布配置chan
	PubConfigChannel chan<- model.Config = configChannel
)

func init() {
	pubConfig()
}

func pubConfig() {
	go func() {
		for {
			conf := <-configChannel
			logger.Debugf("发布App：%d，配置：%s-%v by 模式：%s", conf.AppID, conf.Name, conf.Value, conf.Mode)
		}
	}()
}
