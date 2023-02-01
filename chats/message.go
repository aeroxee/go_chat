package chats

import "time"

type Message struct {
	UserID    uint      `json:"user_id"`
	GroupID   uint      `json:"group_id"`
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}

func newMessage() *Message {
	return &Message{}
}
