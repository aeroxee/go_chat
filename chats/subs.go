package chats

import (
	"github.com/gorilla/websocket"
	"log"
	"time"
)

const (
	writeWait      = time.Second * 10
	pongWait       = time.Second * 60
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

type subscription struct {
	group uint
	conn  *connection
	hub   *Hub
}

func (s *subscription) readPump() {
	c := s.conn
	defer func() {
		c.ws.Close()
		s.hub.unregister <- s
	}()
	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error {
		c.ws.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	for {
		message := newMessage()
		err := c.ws.ReadJSON(message)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message.Timestamp = time.Now()
		message.GroupID = s.group
		s.hub.broadcast <- message
	}
}

func (s *subscription) writePump() {
	c := s.conn
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.ws.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.ws.WriteJSON(message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.ws.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
