package web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	ngin "github.com/nilorg/pkg/gin"
	"github.com/nilorg/pkg/gin/route"
	"github.com/spf13/viper"
)

// HomeController 首页控制器
type HomeController struct {
}

// NewHomeController ...
func NewHomeController() *HomeController {
	return &HomeController{}
}

// Route ... 路由
func (home *HomeController) Route() []route.Route {
	return []route.Route{
		{
			Name:         "首页",
			Method:       http.MethodGet,
			RelativePath: "/",
			Auth:         true,
			HandlerFunc:  ngin.WebControllerFunc(home.Index, "index"),
		},
		{
			Name:         "首页",
			Method:       http.MethodGet,
			RelativePath: "/index.html",
			Auth:         true,
			HandlerFunc:  ngin.WebControllerFunc(home.Index, "index"),
		},
	}
}

// Index ...
func (*HomeController) Index(ctx *ngin.WebContext) {
	ctx.RenderPage(gin.H{
		"title": "index...",
	})
}

// GetWebInfo ...
func (*HomeController) GetWebInfo(name string) interface{} {
	return viper.Get(fmt.Sprintf("web.info.%s", name))
}
