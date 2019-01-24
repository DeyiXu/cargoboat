package web

import (
	"github.com/gin-gonic/gin"
	ngin "github.com/nilorg/pkg/gin"
)

// ErrorController ...
type ErrorController struct {
}

// NewErrorController ...
func NewErrorController() *ErrorController {
	return &ErrorController{}
}

// Error404 ...
func (*ErrorController) Error404(ctx *ngin.WebContext) {
	if ctx.GetCurrentAccount() == nil {
		ctx.HTML(404, "error_404.tmpl", nil)
	} else {
		ctx.RenderPage(gin.H{
			"title": "404",
		})
	}
}
