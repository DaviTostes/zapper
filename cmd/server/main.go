package main

import (
	"gozapper/internal/handler"
	"gozapper/internal/whatsapp"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	whatsapp.ConnectToWhatsApp()

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
	hxGroup.GET("/panel-new-contact", handler.HandleNewContact)

	router.Run(":" + os.Getenv("PORT"))
}
