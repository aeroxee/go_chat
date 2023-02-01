package chats

type Hub struct {
	groups     map[uint]map[*connection]bool
	broadcast  chan *Message
	register   chan *subscription
	unregister chan *subscription
}

func NewHub() *Hub {
	return &Hub{
		groups:     make(map[uint]map[*connection]bool),
		register:   make(chan *subscription),
		unregister: make(chan *subscription),
		broadcast:  make(chan *Message),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case s := <-h.register:
			connections := h.groups[s.group]
			if connections == nil {
				connections = make(map[*connection]bool)
				h.groups[s.group] = connections
			}
			h.groups[s.group][s.conn] = true
		case s := <-h.unregister:
			connections := h.groups[s.group]
			if connections != nil {
				if _, ok := connections[s.conn]; ok {
					delete(connections, s.conn)
					close(s.conn.send)
					if len(connections) == 0 {
						delete(h.groups, s.group)
					}
				}
			}
		case m := <-h.broadcast:
			connections := h.groups[m.GroupID]
			for c := range connections {
				select {
				case c.send <- m:
				default:
					close(c.send)
					delete(connections, c)
					if len(connections) == 0 {
						delete(h.groups, m.GroupID)
					}
				}
			}
		}
	}
}
