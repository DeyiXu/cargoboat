package main

import (
	"log"
	"os"
	"os/signal"

	// 加载配置文件
	"github.com/cargoboat/cargoboat/model"
	_ "github.com/cargoboat/cargoboat/module/config"
	_ "github.com/cargoboat/cargoboat/module/errors"

	// 初始化存储
	"github.com/cargoboat/cargoboat/module/store"
	"github.com/cargoboat/cargoboat/server"
)

func init() {
	store.Start()
	model.AutoMigrate()
	server.Start()
}
func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	defer close()
	log.Println("Server exiting")
}

func close() {
	server.Close()
	store.Close()
}
