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

var Timer uint

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Не удалось загрузить .env")
	}

	token := os.Getenv("TOKEN")
	if token == "" {
		log.Fatal("TOKEN не задан")
	}

	err = db.InitDB("../../.links.db")
	if err != nil {
		log.Fatal("Ошибка инициализации БД:", err)
	}

	b, err := tele.NewBot(tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 30 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
	}

	bot.Setup(b)

	go func() {
		ticker := time.NewTicker(1*time.Hour)
		for range ticker.C {
			url, err := db.RandomLink(bot.UserID)
			if err == nil {
				recipient := &tele.Chat{ID:	bot.UserID}
				_, err := b.Send(recipient, url)
				if err != nil {
					log.Fatal(err)
				}
			} else {
				log.Println("Ошибка получения случайной ссылки:", err)
			}
		}
	}()

	log.Println("Бот запущен...")
	b.Start()
}