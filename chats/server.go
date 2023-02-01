package chats

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ServeWS(hub *Hub, groupID uint, w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	c := &connection{ws: ws, send: make(chan *Message, 255)}
	s := &subscription{
		conn:  c,
		hub:   hub,
		group: groupID,
	}
	s.hub.register <- s

	go s.writePump()
	go s.readPump()
}
