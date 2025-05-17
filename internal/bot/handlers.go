package bot

import (
	tele "gopkg.in/telebot.v4"
	//"log"
	"regexp"
	"tbot/internal/db"
)

var UserID int64
var urlRegex = regexp.MustCompile(`(?i)\b((https?|ftp):\/\/[^\s/$.?#].[^\s]*)`)

func isURL(msg string) bool {
	return urlRegex.MatchString(msg)
}

func Setup(b *tele.Bot) {
	b.Handle(tele.OnText, func(c tele.Context) error {
		UserID = c.Sender().ID
		msg := c.Text()

		if isURL(msg) {
			db.AddLink(UserID, msg)
			return c.Send("Материал сохранён")
		} else {
			return c.Send("Нужно использовать ссылку")
		}
	})
}