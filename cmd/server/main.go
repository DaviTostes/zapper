package main

import (
	"gozapper/internal/handler"
	"gozapper/internal/whatsapp"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	whatsapp.ConnectToWhatsApp()

	router := gin.Default()

	router.SetFuncMap(template.FuncMap{
		"formatDate": func(date time.Time) string { return date.Format("02/01 15:04") },
		"formatName": func(name string) string {
			splited := strings.Split(name, " ")
			if len(splited) == 1 {
				return string(splited[0][0])
			}
			return string(splited[0][0]) + string(splited[1][0])
		},
	})

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
	hxGroup.GET("/panel-cards", handler.HandlePanelCards)
	hxGroup.GET("/panel-details", handler.HandlePanelDetails)
	hxGroup.GET("/panel-new-contact", handler.HandleNewContact)

	router.Run(":" + os.Getenv("PORT"))
}
