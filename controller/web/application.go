package web

import (
	"net/http"

	"github.com/nilorg/pkg/gin/route"

	"github.com/gin-gonic/gin"
	ngin "github.com/nilorg/pkg/gin"
)

// ApplicationController 应用程序控制器
type ApplicationController struct {
}

// NewApplicationController ...
func NewApplicationController() *ApplicationController {
	return &ApplicationController{}
}

// Route ... 路由
func (app *ApplicationController) Route() []route.Route {
	return []route.Route{
		{
			Name:         "应用程序列表",
			Method:       http.MethodGet,
			RelativePath: "/application/list",
			Auth:         true,
			HandlerFunc:  ngin.WebControllerFunc(app.List, "applicationList"),
		},
	}
}

// List ...
func (*ApplicationController) List(ctx *ngin.WebContext) {
	ctx.RenderPage(gin.H{
		"title": "apps list...",
	})
}
