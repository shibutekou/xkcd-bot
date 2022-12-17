package main

import (
	"fmt"
	"github.com/bruma1994/xkcd-bot/internal/xkcd"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TGTOKEN"))
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "help":
			msg.Text = fmt.Sprintf("Отправь мне номер комикса, который хочешь прочитать :D. Наример: /%d",
				rand.Intn(2712))
		case "status":
			msg.Text = "Я в порядке :3"
		default:
			if number, err := strconv.Atoi(strings.TrimPrefix(update.Message.Text, "/")); err == nil {
				if number > 0 && number <= 2712 {
					comics := xkcd.GetComics(update.Message.Text)
					msg.Text = fmt.Sprintf("Название: %s\n%v",
						comics.Title,
						comics.Img)
				} else {
					msg.Text = "Такого комикса нет :("
				}
			} else {
				msg.Text = "Я не знаю такой команды ¯\\_(ツ)_/¯"
			}
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
