package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"tbot/internal/bot"
	"tbot/internal/db"

	tele "gopkg.in/telebot.v4"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Не удалось загрузить .env")
	}

	token := os.Getenv("TOKEN")
	if token == "" {
		log.Fatal("TOKEN не задан")
	}

	db.InitDB("../../.links.db")

	b, err := tele.NewBot(tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 30 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
	}

	bot.Setup(b)

	go func() {
		ticker := time.NewTicker(3*time.Hour)
		for range ticker.C {
			url := db.RandomLink(bot.UserID)
			recipient := &tele.Chat{ID:	bot.UserID}
			_, err := b.Send(recipient, url)
			if err != nil {
				log.Fatal(err)
			}
		}
	}()

	log.Println("Бот запущен...")
	b.Start()
}