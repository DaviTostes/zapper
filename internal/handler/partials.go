package handler

import (
	"gozapper/internal/whatsapp"
	"net/http"
	"time"

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

type Conversation struct {
	Number string
	Name   string
	Status string
}

func HandlePanelCards(ctx *gin.Context) {
	data := struct {
		Conversations []Conversation
	}{
		Conversations: []Conversation{
			{Number: "22992594565", Name: "Davi Tostes", Status: "DONE"},
			{Number: "22988384972", Name: "Maria Eduarda", Status: "PENDING"},
			{Number: "22992594556", Name: "Andrea Siqueira", Status: "CANCELLED"},
			{Number: "21929293233", Name: "Jose Das Couves Almeida", Status: "TESTE"},
		},
	}

	ctx.HTML(http.StatusOK, "panel-cards.html", data)
}

type Client struct {
	Number string
	Name   string
	Status string
}

type Message struct {
	Sender    string
	Content   string
	Timestamp time.Time
}

func HandlePanelDetails(ctx *gin.Context) {
	data := struct {
		Client   Client
		Messages []Message
	}{
		Client: Client{Name: "Davi Tostes", Number: "2299259465", Status: "DONE"},
		Messages: []Message{
			{Sender: "bot", Content: "Ola, tudo bem?", Timestamp: time.Now()},
			{Sender: "user", Content: "Tudo e vc?", Timestamp: time.Now()},
			{Sender: "bot", Content: "dibas", Timestamp: time.Now()},
			{Sender: "user", Content: "Tudo e vc?", Timestamp: time.Now()},
			{Sender: "bot", Content: "dibas", Timestamp: time.Now()},
			{Sender: "user", Content: "Tudo e vc?", Timestamp: time.Now()},
			{Sender: "bot", Content: "dibas", Timestamp: time.Now()},
			{Sender: "user", Content: "Tudo e vc?", Timestamp: time.Now()},
			{Sender: "bot", Content: "dibas", Timestamp: time.Now()},
			{Sender: "user", Content: "Tudo e vc?", Timestamp: time.Now()},
			{Sender: "bot", Content: "dibas", Timestamp: time.Now()},
			{Sender: "user", Content: "Tudo e vc?", Timestamp: time.Now()},
			{Sender: "bot", Content: "dibas", Timestamp: time.Now()},
			{Sender: "user", Content: "Tudo e vc?", Timestamp: time.Now()},
			{Sender: "bot", Content: "dibas", Timestamp: time.Now()},
		},
	}

	ctx.HTML(http.StatusOK, "panel-details.html", data)
}

func HandleNewContact(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "panel-new-contact.html", nil)
}
