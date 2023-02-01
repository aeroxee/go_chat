package chats

import "github.com/gorilla/websocket"

type connection struct {
	ws   *websocket.Conn
	send chan *Message
}
