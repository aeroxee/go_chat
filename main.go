package main

import (
	"github.com/aZ4ziL/go_chat/chats"
	"github.com/aZ4ziL/go_chat/handlers"
	"github.com/aZ4ziL/go_chat/middlewares"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func main() {
	r := mux.NewRouter()
	// static
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	r.PathPrefix("/media/").Handler(http.StripPrefix("/media/", http.FileServer(http.Dir("./media"))))

	// Global Middleware
	r.Use(middlewares.LogMiddleware)

	// webSocket
	hub := chats.NewHub()
	go hub.Run()
	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		roomID, _ := strconv.Atoi(query.Get("room_id"))
		chats.ServeWS(hub, uint(roomID), w, r)
	})

	r.StrictSlash(true).HandleFunc("/login/", handlers.Login)
	r.StrictSlash(true).HandleFunc("/register/", handlers.Register)

	// Home
	r.StrictSlash(true).HandleFunc("/", handlers.Home)

	// chat
	r.StrictSlash(true).HandleFunc("/chat", handlers.ChatHandler)

	http.ListenAndServe("localhost:8000", r)
}
