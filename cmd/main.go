package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/cargoboat/cargoboat"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"

	"github.com/nilorg/pkg/logger"

	// 加载配置文件
	_ "github.com/cargoboat/cargoboat/module/config"
)

var (
	c *cargoboat.Cargoboat
)

func main() {

	logger.Debugln("启动成功...")

	var err error
	redisClient := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	_, err = redisClient.Ping().Result()
	if err != nil {
		logger.Fatalln(err)
	}
	c, err = cargoboat.NewCargoboat(viper.GetString("system.config_dir"), redisClient)
	if err != nil {
		logger.Fatalln(err)
	}
	err = c.Start()
	if err != nil {
		logger.Fatalln(err)
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	defer close()
	log.Println("Server exiting")
}

func close() {
	c.Stop()
}
