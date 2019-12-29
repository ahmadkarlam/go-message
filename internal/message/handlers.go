package message

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ahmadkarlam/go-message/internal/message/request"
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

func (h *Handler) Store(c *gin.Context) {
	var req request.StoreMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	message := h.Service.AddMessage(req.Body)

	if err := websocket.GetWebsocket().Broadcast(message.Body); err != nil {
		log.Println("error broadcast: " + err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "success add new message",
		"data":    message,
	})
}
