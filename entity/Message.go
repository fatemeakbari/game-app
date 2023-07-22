package entity

import "time"

type Message struct {
	ID         uint
	Content    string
	CreateDate time.Time

	RepliedMessageId uint
	UserID           uint
	GroupID          uint
}
