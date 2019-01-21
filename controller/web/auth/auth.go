package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ngin "github.com/nilorg/pkg/gin"
	"github.com/nilorg/pkg/logger"
)

// GetLogin ...
func GetLogin(ctx *ngin.WebContext) {
	redirectURL := ctx.Query("redirect_url")
	logger.Debugf("GetLogin redirectURL:%s", redirectURL)
	ctx.RenderSinglePage(gin.H{
		"title":        "Login...",
		"redirect_url": redirectURL,
	})
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
