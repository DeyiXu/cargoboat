package server

import (
	"time"

	"github.com/appleboy/gin-jwt"
	"github.com/cargoboat/cargoboat/controller/application"
	"github.com/cargoboat/cargoboat/controller/auth"
	"github.com/cargoboat/cargoboat/controller/config"
	"github.com/gin-gonic/gin"
	ngin "github.com/nilorg/pkg/gin"
)

var (
	// the jwt middleware
	authMiddleware = &jwt.GinJWTMiddleware{
		Realm:         "test zone",
		Key:           []byte("secret key"),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		Authenticator: auth.Authenticator,
		Authorizator:  auth.Authorizator,
		Unauthorized:  auth.Unauthorized,
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	}
)

// setRouter 设置路由
func setRouter(handler *gin.Engine) {
	handler.GET("/", func(c *gin.Context) {
		c.String(200, "welcome cogo server")
	})

	handler.POST("/login", authMiddleware.LoginHandler)
	auth := handler.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		// auth.GET("/hello", helloHandler)
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)

		auth.GET("/applications", ngin.WebAPIControllerFunc(application.Get))
		auth.GET("/applications/:id", ngin.WebAPIControllerFunc(application.GetOne))
		auth.GET("/applications/:id/configs", ngin.WebAPIControllerFunc(application.GetConfigs))
		auth.POST("/applications", ngin.WebAPIControllerFunc(application.Post))
		auth.DELETE("/applications/:id", ngin.WebAPIControllerFunc(application.Delete))

		auth.GET("/configs", ngin.WebAPIControllerFunc(config.Get))
		auth.POST("/configs", ngin.WebAPIControllerFunc(config.Post))
		auth.PUT("/configs/:id", ngin.WebAPIControllerFunc(config.Put))
		auth.DELETE("/configs/:id", ngin.WebAPIControllerFunc(config.Delete))
	}
}
