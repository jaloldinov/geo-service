package api

import (
	"geo/api/handler"

	"github.com/gin-gonic/gin"
)

func NewServer(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	// authentication sign up and login
	r.POST("/auth/login", h.Login)
	r.POST("/auth/signup", h.SignUp)

	r.POST("/ws/createRoom", h.CreateRoom)
	r.GET("/ws/joinRoom/:roomId", h.JoinRoom)
	r.GET("/ws/getRooms", h.GetRooms)
	r.GET("/ws/getClients/:roomId", h.GetClients)

	//r.POST("ws/file-upload", h.HandleFileTransfer)

	return r
}
