package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cargoboat/cargoboat/bll"
	cerrors "github.com/cargoboat/cargoboat/module/errors"
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
		{
			Name:         "应用程序添加模式",
			Method:       http.MethodPost,
			RelativePath: "/application/mode/add",
			Auth:         true,
			HandlerFunc:  ngin.WebAPIControllerFunc(app.AddMode),
		},
	}
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

type addModeModel struct {
	ModeName string `json:"mode_name"`
	AppID    int64  `json:"app_id"`
}

// AddMode 添加模式
func (*ApplicationController) AddMode(ctx *ngin.WebAPIContext) {
	model := addModeModel{}
	if err := ctx.Bind(&model); err != nil {
		ctx.ResultError(cerrors.GetBusinessError(1001))
		return
	}
	modeID, err := bll.Application.AddMode(model.AppID, model.ModeName)
	if err != nil {
		ctx.ResultError(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"mode_id": modeID,
	})
}
