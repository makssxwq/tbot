package db

import "log"

func AddLink(userID int64, url string) {
	_, err := DB.Exec("INSERT INTO links (user_id, url) VALUES (?, ?)", userID, url)
	log.Printf("Ссылка успешно добавлена в бд - USERID: %v, URL: %v", userID, url)
	if err != nil {
		log.Fatal(err)
	}
}

func RandomLink(userID int64) (string) {
	row := DB.QueryRow(`SELECT url FROM links WHERE user_id = ? ORDER BY RANDOM() LIMIT 1`, userID)
	var url string
	err := row.Scan(&url)
	log.Printf("Выбрана ссылка для пользователя: %v, URL: %v", userID, url)
	if err != nil {
		log.Fatal(err)
	}
	return url
}