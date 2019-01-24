package web

import (
	"github.com/nilorg/pkg/gin/route"
)

var (
	Application = NewApplicationController()
	Auth        = NewAuthController()
	Error       = NewErrorController()
	Home        = NewHomeController()
)

func Router() []route.Router {
	routerList := make([]route.Router, 0)
	routerList = append(routerList, Application)
	routerList = append(routerList, Auth)
	routerList = append(routerList, Home)
	return routerList
}
