package bot

import (
	tele "gopkg.in/telebot.v4"
	"log"
	"tbot/internal/db"
)

func Setup(b *tele.Bot) {
	b.Handle(tele.OnText, func(c tele.Context) error {
		userID := c.Sender().ID
		msg := c.Text()

		err := db.AddLink(userID, msg)
		if err != nil {
			log.Println("Ошибка при сохранении:", err)
			return c.Send("Не удалось сохранить")
		}

		return c.Send("Материал сохранён")
	})
}