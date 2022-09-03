package main

import (
	"log"
	"fmt"
	"github.com/aaletov/go-chat/pkg/client"
)

func main() {
	log.Println("Starting client")

	var c client.Client
	var command string
	for {
		_, err := fmt.Scan(&command)
		if err != nil {
			log.Printf("Scan error: %v", err)
			break
		}

		switch command {
		case "/genkey":
			err = c.GenKey()
		case "/usekey":
			err = c.UseKey()
		case "/getkey":
			err = c.GetKey()
		case "/chat":
			err = c.Chat()
		default:
			fmt.Println("No such command")
		}

		if err != nil {
			log.Printf("Command error: %v\n", err)
			break
		}
	}

	log.Println("Exiting client..")
}