package whatsapp

import (
	"context"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
)

var (
	client *whatsmeow.Client
	ctx    = context.Background()
)

func IsConnected() bool {
	return client.IsConnected()
}

func GetPushName() string {
	return client.Store.PushName
}

func ConnectToWhatsApp() (string, error) {
	dbLog := waLog.Stdout("Database", "INFO", true)
	container, err := sqlstore.New(ctx, "sqlite3", "file:db/store.db?_foreign_keys=on", dbLog)
	if err != nil {
		return "", err
	}

	deviceStore, err := container.GetFirstDevice(ctx)
	if err != nil {
		return "", err
	}

	clientLog := waLog.Stdout("Client", "INFO", true)
	client = whatsmeow.NewClient(deviceStore, clientLog)

	client.AddEventHandler(eventHandler)

	if client.Store.ID != nil {
		err = client.Connect()
		if err != nil {
			return "", err
		}

		return "", nil
	}

	qrChan, _ := client.GetQRChannel(ctx)
	err = client.Connect()
	if err != nil {
		return "", err
	}

	for evt := range qrChan {
		if evt.Event != "code" {
			continue
		}

		return evt.Code, nil
	}

	return "", nil
}

func SendMessage(recipientJID string, message string) {
	jid, err := types.ParseJID(recipientJID)
	if err != nil {
		fmt.Printf("Error parsing JID: %v\n", err)
		return
	}

	msg := &waE2E.Message{
		Conversation: &message,
	}

	_, err = client.SendMessage(ctx, jid, msg)
	if err != nil {
		fmt.Printf("Error sending message: %v\n", err)
	} else {
		fmt.Println("Message sent successfully!")
	}
}

func eventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Connected:
		fmt.Println("✅ Connection successful")

	case *events.Message:
		sender := v.Info.Sender.String()
		groupName := ""
		if v.Info.IsGroup {
			groupInfo, err := client.GetGroupInfo(v.Info.Chat)
			if err != nil {
				groupName = "an unknown group"
			} else {
				groupName = "'" + groupInfo.Name + "'"
			}
			fmt.Printf("Received a group message in %s from %s\n", groupName, sender)
		} else {
			fmt.Printf("Received a direct message from %s\n", sender)
		}

		messageText := ""
		if msg := v.Message.GetConversation(); msg != "" {
			messageText = msg
		} else if extMsg := v.Message.GetExtendedTextMessage(); extMsg != nil {
			messageText = extMsg.GetText()
		} else if imgMsg := v.Message.GetImageMessage(); imgMsg != nil {
			messageText = imgMsg.GetCaption()
			if messageText != "" {
				messageText = "[Image] " + messageText
			} else {
				messageText = "[Image]"
			}
		} else {
			messageText = "[Unsupported message type]"
		}

		fmt.Printf("  -> Message: %s\n", messageText)

		if strings.ToLower(messageText) == "ping" {
			fmt.Println("  -> Replying with 'pong'")
			SendMessage(sender, "pong")
		}

	case *events.Receipt:
		fmt.Printf("Received a receipt for message %s from %s: %s\n", v.MessageIDs, v.MessageSource, v.Type)

	case *events.Disconnected:
		fmt.Println("❌ Disconnected from WhatsApp")
	}
}
