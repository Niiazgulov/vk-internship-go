package main

import (
	"log"

	"github.com/Niiazgulov/vk-internship-go.git/cmd/client"
)

func main() {
	offset := 0
	for {
		updates, err := client.GetUpdates(offset)
		if err != nil {
			log.Fatal(err)
		}
		for _, upd := range updates {
			client.SendMessages(upd)
			offset = upd.UpdateID + 1
		}
	}
}
