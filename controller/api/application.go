package api

import (
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
	return []route.Route{}
}
