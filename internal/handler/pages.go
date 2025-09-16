package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleAdmin(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin.html", nil)
}
