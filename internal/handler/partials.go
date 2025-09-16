package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleSidebar(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "sidebar.html", nil)
}

func HandlePanel(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "panel.html", nil)
}

func HandleNewInstance(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "new-instance.html", nil)
}
