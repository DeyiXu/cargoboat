package home

import (
	"fmt"

	"github.com/gin-gonic/gin"
	ngin "github.com/nilorg/pkg/gin"
	"github.com/spf13/viper"
)

// Index ...
func Index(ctx *ngin.WebContext) {
	ctx.RenderPage(gin.H{
		"title": "index...",
	})
}

// GetWebInfo ...
func GetWebInfo(name string) interface{} {
	return viper.Get(fmt.Sprintf("web.info.%s", name))
}
