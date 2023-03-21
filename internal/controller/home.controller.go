package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeController struct {
}

func InitHomeController(r *gin.RouterGroup) {
	controller := &HomeController{}

	r.GET("", controller.Home)
}

func (h HomeController) Home(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title": "xxx",
	})
}
