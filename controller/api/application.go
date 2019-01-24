package api

import (
	"net/http"

	"github.com/cargoboat/cargoboat/bll"

	ngin "github.com/nilorg/pkg/gin"
	nginBinding "github.com/nilorg/pkg/gin/binding"
	nginModel "github.com/nilorg/pkg/gin/model"
	"github.com/nilorg/pkg/gin/route"
)

// ApplicationController 应用程序控制器
type ApplicationController struct {
}

// NewApplicationController ...
func NewApplicationController() *ApplicationController {
	return &ApplicationController{}
}

// List ...
func (*ApplicationController) List(ctx *ngin.WebAPIContext) {
	model := nginModel.JQueryDataTablesParameters{}
	result := nginModel.JQueryDataTableResult{
		Data: []interface{}{},
	}
	if err := ctx.MustBindWith(&model, nginBinding.JQueryDataTables); err != nil {
		result.Error = err.Error()
	} else {
		apps, total, verr := bll.Application.QueryPage(model.Start, model.Length, "")
		if verr != nil {
			result.Error = verr.Error()
		} else {
			result.RecordsTotal = total
			result.RecordsFiltered = total
			result.Data = apps
		}
	}
	result.Draw = model.Draw
	ctx.JSON(http.StatusOK, result)
}

// Route ... 路由
func (app *ApplicationController) Route() []route.Route {
	return []route.Route{
		{
			Name:         "应用程序列表",
			Method:       http.MethodGet,
			RelativePath: "/application/list",
			Auth:         true,
			HandlerFunc:  ngin.WebAPIControllerFunc(app.List),
		},
	}
}
