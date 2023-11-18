package handler

import (
	"geo/models"
	"geo/pkg/helper"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type CreateRoomReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) CreateRoom(c *gin.Context) {
	var req CreateRoomReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.hub.Rooms[req.ID] = &Room{
		ID:      req.ID,
		Name:    req.Name,
		Clients: make(map[string]*Client),
	}

	c.JSON(http.StatusOK, req)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) HandleFileTransfer(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Read the uploaded file content
	//fileContent, err := file.Open()
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	//defer fileContent.Close()
	path, _ := helper.Service{}.Upload(context.Background(), file, "file/")
	//fmt.Println(path)
	//content, err := io.ReadAll(fileContent)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}

	roomd_id := c.Query("room_id")

	m := &File{
		Name:    file.Filename,
		Content: path,
		RoomID:  roomd_id,
	}

	h.hub.SendFile <- m

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})

}

func (h *Handler) JoinRoom(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	roomID := c.Param("roomId")
	clientID := c.Query("userId")
	username := c.Query("username")

	cl := &Client{
		Conn:     conn,
		Message:  make(chan *models.Message, 10),
		File:     make(chan *File, 10),
		ID:       clientID,
		RoomID:   roomID,
		Username: username,
	}

	m := &models.Message{
		Content:  "A new user has joined the room",
		RoomID:   roomID,
		Username: username,
	}

	h.hub.Register <- cl
	h.hub.Broadcast <- m

	go cl.writeMessage()
	cl.readMessage(h.hub)
}

type RoomRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) GetRooms(c *gin.Context) {
	rooms := make([]RoomRes, 0)

	for _, r := range h.hub.Rooms {
		rooms = append(rooms, RoomRes{
			ID:   r.ID,
			Name: r.Name,
		})
	}

	c.JSON(http.StatusOK, rooms)
}

type ClientRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (h *Handler) GetClients(c *gin.Context) {
	var clients []ClientRes
	roomId := c.Param("roomId")

	if _, ok := h.hub.Rooms[roomId]; !ok {
		clients = make([]ClientRes, 0)
		c.JSON(http.StatusOK, clients)
	}

	for _, c := range h.hub.Rooms[roomId].Clients {
		clients = append(clients, ClientRes{
			ID:       c.ID,
			Username: c.Username,
		})
	}

	c.JSON(http.StatusOK, clients)
}
