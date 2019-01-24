package api

import "github.com/nilorg/pkg/gin/route"

var (
	Application = NewApplicationController()
	Auth        = NewAuthController()
)

func Router() []route.Router {
	routerList := make([]route.Router, 0)
	routerList = append(routerList, Application)
	routerList = append(routerList, Auth)
	return routerList
}
