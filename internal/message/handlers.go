package message

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ahmadkarlam/go-message/internal/message/service"
	"github.com/ahmadkarlam/go-message/pkg/websocket"
)

type Handler struct {
	Service service.MessageService
}

func NewHandler() Handler {
	return Handler{
		Service: service.NewMessageService(),
	}
}

func (h *Handler) Get(c *gin.Context) {
	messages := h.Service.GetAllMessage()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "success get all message",
		"data":    messages,
	})
}

func (h *Handler) Add(c *gin.Context) {
	message := h.Service.AddMessage(c.PostForm("body"))

	if err := websocket.GetWebsocket().Broadcast(message.Body); err != nil {
		log.Println("error broadcast: " + err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "success add new message",
		"data":    message,
	})
}
