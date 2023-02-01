package models

import "time"

type Group struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	Members     []*User   `gorm:"many2many:group_members" json:"members,omitempty"`
	Title       string    `gorm:"size:30" json:"title"`
	Logo        string    `gorm:"size:255;null" json:"logo"`
	Description string    `gorm:"type:text" json:"description"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
	Chats       []Chat    `gorm:"foreignKey:GroupID"`
}

func (Group) TableName() string {
	return "groups"
}

// CreateNewGroup create new group
func CreateNewGroup(group *Group) error {
	return db.Create(group).Error
}

// GetGroupByID get group by id
func GetGroupByID(id uint) (Group, error) {
	var group Group
	err := db.Model(&Group{}).Where("id = ?", id).Preload("Members").Preload("Chats").
		First(&group).Error
	return group, err
}

// GetAllGroups get all groups
func GetAllGroups() []Group {
	var groups []Group
	db.Model(&Group{}).Preload("Chats").Preload("Members").Find(&groups)
	return groups
}
