package bot

import (
	tele "gopkg.in/telebot.v4"
	"log"
	"tbot/internal/db"
)

var UserID int64

func Setup(b *tele.Bot) {
	b.Handle(tele.OnText, func(c tele.Context) error {
		UserID = c.Sender().ID
		msg := c.Text()

		err := db.AddLink(UserID, msg)
		if err != nil {
			log.Println("Ошибка при сохранении:", err)
			return c.Send("Не удалось сохранить")
		}

		return c.Send("Материал сохранён")
	})

	/*b.Handle("/timer", func(c tele.Context) error {
		var err error
		//args := c.Args()

		return err
	})*/
}