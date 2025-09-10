package main

import (
	"gozapper/internal/whatsmeow"
	"os"
)

func main() {
	args := os.Args
	switch args[1] {
	case "connect":
		whatsmeow.ConnectClient()
	}
}
