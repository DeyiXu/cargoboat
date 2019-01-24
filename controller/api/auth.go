package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ngin "github.com/nilorg/pkg/gin"
	"github.com/nilorg/pkg/gin/route"
	"github.com/nilorg/pkg/logger"
)

// AuthController 授权控制器
type AuthController struct {
}

// NewAuthController ...
func NewAuthController() *AuthController {
	return &AuthController{}
}

// Route ... 路由
func (auth *AuthController) Route() []route.Route {
	return []route.Route{
		{
			Name:         "登录",
			Method:       http.MethodPost,
			RelativePath: "/login",
			Auth:         false,
			HandlerFunc:  ngin.WebAPIControllerFunc(auth.PostLogin),
		},
	}
}

// PostLogin ...
func (*AuthController) PostLogin(ctx *ngin.WebAPIContext) {
	userName := ctx.PostForm("username")
	pass := ctx.PostForm("pass")
	me := ctx.PostForm("me")
	redirectURL := ctx.PostForm("redirect_url")
	logger.Debugf("name:%s|pass:%s|me:%s", userName, pass, me)
	if userName == "cargoboat" && pass == "cargoboat" {
		ctx.SetCurrentAccount(gin.H{
			"username": userName,
			"name":     "德意洋洋",
			"status":   "在线",
			"cover":    "/assets/img/IMG_3476.jpg",
		})
		if redirectURL == "" {
			redirectURL = "/index.html"
		}
		// post redirect get request by http status code 303.
		// ctx.Redirect(http.StatusSeeOther, redirectURL)
		ctx.Status(200)
	} else {
		// ctx.Redirect(http.StatusSeeOther, "/login.html")
		ctx.JSON(400, gin.H{
			"error": "账号密码错误",
		})
	}
}

// GetMenuData 获取菜单数据
func (*AuthController) GetMenuData(value interface{}) gin.H {
	logger.Debugln("getMenuData...")
	return gin.H{
		"account": value,
		"info":    "...",
	}
}

// GetNavigationData 获取导航数据
func (*AuthController) GetNavigationData(value interface{}) gin.H {
	logger.Debugln("GetNavigationData...")
	return gin.H{
		"account": value,
		"info":    "navigation...",
	}
}
