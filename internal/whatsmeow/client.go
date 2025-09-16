package whatsmeow

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"

	_ "github.com/mattn/go-sqlite3"
)

func eventHandler(evt any) {
	switch v := evt.(type) {
	case *events.Message:
		fmt.Println("Received Message: ", v.Info.Sender)
	}
}

var (
	ctx = context.Background()
)

func createLogDB() (*sqlstore.Container, error) {
	dbLog := waLog.Stdout("Database", "", true)
	container, err := sqlstore.New(ctx, "sqlite3", "file:store.db?_foreign_keys=on", dbLog)
	return container, err
}

func ConnectClient() error {
	container, err := createLogDB()
	if err != nil {
		return err
	}

	deviceStore, err := container.GetFirstDevice(ctx)
	if err != nil {
		return err
	}

	clientLog := waLog.Stdout("Client", "", true)
	client := whatsmeow.NewClient(deviceStore, clientLog)
	client.AddEventHandler(eventHandler)

	if client.Store.ID != nil {
		err = client.Connect()
		if err != nil {
			return err
		}
	}

	qrChan, _ := client.GetQRChannel(context.Background())
	err = client.Connect()
	if err != nil {
		return err
	}

	for evt := range qrChan {
		if evt.Event == "code" {
		} else {
			fmt.Println("Login event:", evt.Event)
		}
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	client.Disconnect()

	return nil
}
