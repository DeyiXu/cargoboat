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

// MenuRoot 菜单根节点
type MenuRoot struct {
	URL       string
	Name      string
	Icon      string
	MenuItems []*MenuItem
}

// MenuItem 菜单项
type MenuItem struct {
	URL  string
	Name string
	Icon string
}

// GetMenuData 获取菜单数据
func GetMenuData(value interface{}) gin.H {
	logger.Debugln("getMenuData...")
	roots := make([]*MenuRoot, 0)
	roots = append(roots,
		&MenuRoot{
			URL:       "/test",
			Name:      "测试页面",
			Icon:      "fa fa-link",
			MenuItems: nil,
		},
		&MenuRoot{
			URL:  "#",
			Name: "应用程序管理",
			Icon: "fa fa-link",
			MenuItems: []*MenuItem{
				{
					URL:  "/application/list",
					Name: "应用程序列表",
					Icon: "fa fa-circle-o",
				},
			},
		})
	return gin.H{
		"account": value,
		"info":    "...",
		"menus":   roots,
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
