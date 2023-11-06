package main

import (
	"flag"
	"fmt"

	"github.com/uerax/propin/server/core"
	"github.com/uerax/propin/server/pin"
)

var (
	tgApi string
	token string
	chatID string
)

func main() {
	flag.StringVar(&token, "token", "", "TelegramBot的token")
	flag.StringVar(&chatID, "id", "", "接收消息的TelegramId")
	flag.Parse()
	if token == "" || chatID == "" {
		panic("token和chatID不可为空")
	}
	tgApi = fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
	pin.NewPin(120)
	core.Start(tgApi, chatID)
}

