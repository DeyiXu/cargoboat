package server

import (
	authAPIController "github.com/cargoboat/cargoboat/controller/api/auth"

	"github.com/gin-gonic/gin"
	ngin "github.com/nilorg/pkg/gin"
)

// setAPIRouter 设置路由
func setAPIRouter(router *gin.RouterGroup) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	router.POST("/login", ngin.WebControllerFunc(authAPIController.PostLogin, "login"))
}
