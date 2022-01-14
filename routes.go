package main

import (
	"github.com/Strike-official/global-getting-started/internal/controller"
	"github.com/gin-gonic/gin"
)

func routes(router *gin.Engine) {

	ytbot := router.Group("/global-getting-started/")
	{
		ytbot.POST("/getting-started", controller.Getting_started)
	}
}
