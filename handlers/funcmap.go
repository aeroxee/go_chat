package handlers

import (
	"github.com/aZ4ziL/go_chat/models"
	"text/template"
)

var funcMap = template.FuncMap{
	"getUserFullName": getUserFullName,
}

func getUserFullName(id uint) string {
	user, _ := models.GetUserByID(id)
	return user.FirstName + " " + user.LastName
}
