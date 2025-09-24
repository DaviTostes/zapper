package handler

import (
	"gozapper/internal/whatsapp"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleSidebar(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "sidebar.html", nil)
}

func HandlePanel(ctx *gin.Context) {
	data := struct {
		IsConnected  bool
		QRCode       string
		InstanceName string
	}{
		IsConnected:  whatsapp.IsConnected(),
		QRCode:       "",
		InstanceName: "",
	}

	if !data.IsConnected {
		qrcode, err := whatsapp.ConnectToWhatsApp()
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		data.QRCode = qrcode
	} else {
		data.InstanceName = whatsapp.GetPushName()
	}

	ctx.HTML(http.StatusOK, "panel.html", data)
}

func HandlePanelConfig(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "panel-config.html", nil)
}

func HandlePanelAnalysis(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "panel-analysis.html", nil)
}

func HandleNewContact(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "panel-new-contact.html", nil)
}
