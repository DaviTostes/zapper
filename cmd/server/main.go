package main

import (
	"gozapper/internal/handler"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("/static", "web/static")
	router.LoadHTMLGlob("web/templates/**/*.html")

	router.GET("/", handler.HandleAdmin)

	hxGroup := router.Group("/partials").Use(func(ctx *gin.Context) {
		if ctx.GetHeader("HX-Request") != "true" {
			ctx.AbortWithError(http.StatusServiceUnavailable, nil)
			return
		}
	})

	hxGroup.GET("/sidebar", handler.HandleSidebar)
	hxGroup.GET("/panel", handler.HandlePanel)
	hxGroup.GET("/panel-config", handler.HandlePanelConfig)
	hxGroup.GET("/panel-analysis", handler.HandlePanelAnalysis)

	router.Run(":" + os.Getenv("PORT"))
}
