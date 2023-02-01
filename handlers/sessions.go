package handlers

import (
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("goChatSessionID"))
