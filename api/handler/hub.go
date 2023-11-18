package handler

import (
	"geo/models"
)

type Room struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *models.Message
	SendFile   chan *File // New channel for file sending
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client, 5),
		Unregister: make(chan *Client, 5),
		Broadcast:  make(chan *models.Message, 5),
		SendFile:   make(chan *File, 5),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case cl := <-h.Register:
			if _, ok := h.Rooms[cl.RoomID]; ok {
				r := h.Rooms[cl.RoomID]

				if _, ok := r.Clients[cl.ID]; !ok {
					r.Clients[cl.ID] = cl
				}
			}
		case cl := <-h.Unregister:
			if _, ok := h.Rooms[cl.RoomID]; ok {
				if _, ok := h.Rooms[cl.RoomID].Clients[cl.ID]; ok {
					if len(h.Rooms[cl.RoomID].Clients) != 0 {
						h.Broadcast <- &models.Message{
							Content:  "user left the chat",
							RoomID:   cl.RoomID,
							Username: cl.Username,
						}
					}

					delete(h.Rooms[cl.RoomID].Clients, cl.ID)
					close(cl.Message)
					close(cl.File) // Close the file channel
				}
			}

		case m := <-h.Broadcast:
			if _, ok := h.Rooms[m.RoomID]; ok {
				for _, cl := range h.Rooms[m.RoomID].Clients {
					cl.Message <- m
				}
			}

		case f := <-h.SendFile: // Handle file sending
			if _, ok := h.Rooms[f.RoomID]; ok {
				for _, cl := range h.Rooms[f.RoomID].Clients {
					cl.File <- f
				}
			}
		}
	}
}
