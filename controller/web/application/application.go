package application

import (
	"github.com/gin-gonic/gin"
	ngin "github.com/nilorg/pkg/gin"
)

// List ...
func List(ctx *ngin.WebContext) {
	ctx.RenderPage(gin.H{
		"title": "apps list...",
	})
}
