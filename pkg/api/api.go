package api

import (
	"github.com/gin-gonic/gin"

	"github.com/ahmadkarlam/go-message/internal/message"
	"github.com/ahmadkarlam/go-message/internal/websocket"
)

func Start(port string) error {
	r := gin.Default()

	r.GET("/v1/message", message.Get)
	r.POST("/v1/message", message.Add)
	r.POST("/ws", websocket.Initialize)

	err := r.Run(":" + port)
	if err != nil {
		return err
	}

	return nil
}
