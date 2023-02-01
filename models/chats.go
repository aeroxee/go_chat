package models

import "time"

type Chat struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	GroupID   uint      `json:"group_id"`
	Text      string    `gorm:"type:text" json:"text"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func (Chat) TableName() string {
	return "chats"
}

// CreateNewChat function to create new chat to group
func CreateNewChat(chat *Chat) error {
	return db.Create(chat).Error
}

// GetChatByID get chat by id
func GetChatByID(id uint) (Chat, error) {
	var chat Chat
	err := db.Model(&Chat{}).Where("id = ?", id).Find(&chat).Error
	return chat, err
}
