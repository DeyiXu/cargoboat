package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ngin "github.com/nilorg/pkg/gin"
	"github.com/nilorg/pkg/logger"
)

// PostLogin ...
func PostLogin(ctx *ngin.WebContext) {
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

// Logout ...
func Logout(ctx *ngin.WebAPIContext) {
	ctx.DelCurrentAccount()
	ctx.Redirect(http.StatusSeeOther, "/login.html")
}

// GetMenuData 获取菜单数据
func GetMenuData(value interface{}) gin.H {
	logger.Debugln("getMenuData...")
	return gin.H{
		"account": value,
		"info":    "...",
	}
}

// GetNavigationData 获取导航数据
func GetNavigationData(value interface{}) gin.H {
	logger.Debugln("GetNavigationData...")
	return gin.H{
		"account": value,
		"info":    "navigation...",
	}
}
