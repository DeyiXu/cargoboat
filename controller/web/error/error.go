package error

import (
	"github.com/gin-gonic/gin"
	ngin "github.com/nilorg/pkg/gin"
)

// Error404 ...
func Error404(ctx *ngin.WebContext) {
	if ctx.GetCurrentAccount() == nil {
		ctx.HTML(404, "error_404.tmpl", nil)
	} else {
		ctx.RenderPage(gin.H{
			"title": "404",
		})
	}
}
