package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

func GetTokenFlag() string {
	token := ""
	flag.StringVar(&token, "tt", "", "telegram-bot token")
	flag.Parse()
	return token
}

type Token struct {
	TelegramBot string
}

func GetTokenFile() string {
	f, err := os.Open("config.json")
	if err != nil {
		fmt.Println("can't open file: ", err)
	}
	dec := json.NewDecoder(f)
	token := Token{}
	err = dec.Decode(&token)
	if err != nil {
		log.Panic(err)
	}
	return token.TelegramBot
}
