package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoute(r gin.IRouter){

	r.GET("/health-check", func(ctx *gin.Context) {
		result := "My first step :)"
		ctx.JSON(http.StatusOK, result)
	})

}