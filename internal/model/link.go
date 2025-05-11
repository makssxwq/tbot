package model

import "time"

type Link struct {
	ID		int
	userID	int64
	URL 	string
	AddedAt	time.Time
}