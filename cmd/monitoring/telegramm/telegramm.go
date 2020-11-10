package telegramm

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func Send(txt string) {
	bot, err := tgbotapi.NewBotAPI("1379622887:AAH9mv_dtKTqUgrv_myIzc7z5DddSQ_YvOA")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	// Созадаем сообщение
	msg := tgbotapi.NewMessage(-1001418005327, txt)
	// и отправляем его
	bot.Send(msg)
}
