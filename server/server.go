package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/cargoboat/cargoboat/controller/api"
	"github.com/cargoboat/cargoboat/controller/web"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	httpServer *http.Server
)

// Start 启动
func Startup() {
	engine := gin.Default()
	setRouter(engine.Group("/demo"))
	setWeb(engine)
	setWebRouter(engine, web.Router()...)
	setAPIRouter(engine.Group("/api"), api.Router()...)
	httpServer = &http.Server{
		Addr:    viper.GetString("system.addr"),
		Handler: engine,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
}

// Close 关闭
func Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal("HttpServer Shutdown:", err)
	}
}
