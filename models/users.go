package models

import (
	"github.com/aZ4ziL/go_chat/auth"
	"time"
)

type User struct {
	ID           uint       `json:"id"`
	FirstName    string     `gorm:"size:50" json:"first_name"`
	LastName     string     `gorm:"size:50" json:"last_name"`
	Username     string     `gorm:"size:100;uniqueIndex" json:"username"`
	Email        string     `gorm:"size:100;uniqueIndex" json:"email"`
	Password     string     `gorm:"size:128" json:"-"`
	LastLogin    *time.Time `gorm:"null" json:"last_login"`
	DateJoined   time.Time  `gorm:"autoCreateTime" json:"date_joined"`
	GroupAdmins  []Group    `gorm:"foreignKey:UserID" json:"group_admins,omitempty"`
	Chats        []Chat     `gorm:"foreignKey:UserID" json:"chats,omitempty"`
	GroupMembers []*Group   `gorm:"many2many:group_members" json:"group_members,omitempty"`
}

func (User) TableName() string {
	return "users"
}

// CreateNewUser function to create new user
func CreateNewUser(user *User) error {
	user.Password = auth.EncryptionPassword(user.Password)
	return db.Create(user).Error
}

// GetUserByUsername get user by username
func GetUserByUsername(username string) (User, error) {
	var user User
	err := db.Model(&User{}).Where("username = ?", username).Preload("GroupAdmins").
		Preload("Chats").Preload("GroupMembers").First(&user).Error
	return user, err
}

// GetUserByID get user by id
func GetUserByID(id uint) (User, error) {
	var user User
	err := db.Model(&User{}).Where("id = ?", id).Preload("GroupAdmins").
		Preload("Chats").Preload("GroupMembers").First(&user).Error
	return user, err
}
