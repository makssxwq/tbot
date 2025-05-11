package db

import "log"

func AddLink(userID int64, url string) error {
	_, err := DB.Exec("INSERT INTO links (user_id, url) VALUES (?, ?)", userID, url)
	log.Printf("Ссылка успешно добавлена в бд - USERID: %v, URL: %v", userID, url)
	return err
}