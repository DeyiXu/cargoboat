package server

import (
	"time"

	"github.com/nilorg/pkg/gin/route"

	"github.com/appleboy/gin-jwt"
	"github.com/cargoboat/cargoboat/controller/auth"
	"github.com/gin-gonic/gin"
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
func setRouter(handler *gin.RouterGroup) {
	handler.GET("/", func(c *gin.Context) {
		c.String(200, "welcome cargoboat server")
	})

	handler.POST("/login", authMiddleware.LoginHandler)
	auth := handler.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		// auth.GET("/hello", helloHandler)
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)

	}
}

// setAPIRouter 设置路由
func setAPIRouter(api *gin.RouterGroup, controllers ...route.Router) {
	apiAuthRouter := api.Group("/")
	apiAuthRouter.Use(AuthRequired)
	for _, controller := range controllers {
		for _, route := range controller.Route() {
			if route.Auth {
				apiAuthRouter.Handle(route.Method, route.RelativePath, route.HandlerFunc)
			} else {
				api.Handle(route.Method, route.RelativePath, route.HandlerFunc)
			}
		}
	}
}

// setAPIRouter 设置路由
func setWebRouter(web *gin.Engine, controllers ...route.Router) {
	webAuthRouter := web.Group("/")
	webAuthRouter.Use(AuthRequired)
	for _, controller := range controllers {
		for _, route := range controller.Route() {
			if route.Auth {
				webAuthRouter.Handle(route.Method, route.RelativePath, route.HandlerFunc)
			} else {
				web.Handle(route.Method, route.RelativePath, route.HandlerFunc)
			}
		}
	}
}
