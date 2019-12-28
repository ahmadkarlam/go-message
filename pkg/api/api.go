package api

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/ahmadkarlam/go-message/internal/message"
	"github.com/ahmadkarlam/go-message/internal/websocket"
)

func Start(port string) error {
	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("web", false)))
	messageHandler := message.NewHandler()
	r.GET("/v1/message", messageHandler.Get)
	r.POST("/v1/message", messageHandler.Add)
	r.POST("/ws", websocket.Initialize)

	err := r.Run(":" + port)
	if err != nil {
		return err
	}

	return nil
}
