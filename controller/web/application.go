package web

import (
	"net/http"

	"github.com/nilorg/pkg/logger"

	"github.com/cargoboat/cargoboat/bll"

	"github.com/nilorg/pkg/gin/route"
	"github.com/nilorg/sdk/convert"

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
		{
			Name:         "应用程序模式列表",
			Method:       http.MethodGet,
			RelativePath: "/application/mode/list",
			Auth:         true,
			HandlerFunc:  ngin.WebControllerFunc(app.ModeList, "applicationModeList"),
		},
	}
}

// List ...
func (*ApplicationController) List(ctx *ngin.WebContext) {
	ctx.RenderPage(gin.H{
		"title": "apps list...",
	})
}

// ModeList ...
func (*ApplicationController) ModeList(ctx *ngin.WebContext) {
	appID := convert.ToInt64(ctx.Query("app_id"))
	app := bll.Application.GetOneByID(appID)
	logger.Debugln(app)
	modes := bll.Application.GetModeAll(appID)
	ctx.RenderPage(gin.H{
		"app":   app,
		"modes": modes,
	})
}
