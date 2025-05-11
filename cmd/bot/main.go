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

	log.Println("Бот запущен...")
	b.Start()
}