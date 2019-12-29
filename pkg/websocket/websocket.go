package websocket

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var ws *Websocket

// Response that return when web socket write json
type SocketMessage struct {
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

type Websocket struct {
	Upgrader   websocket.Upgrader
	Connection *websocket.Conn
}

func GetWebsocket() *Websocket {
	if ws == nil {
		upgrader := websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		}

		ws = &Websocket{
			Upgrader: upgrader,
		}
	}

	return ws
}

func Run(w http.ResponseWriter, r *http.Request) {
	ws := GetWebsocket()
	conn, err := ws.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("failed to set websocket upgrade: %+v\n", err)
		return
	}
	ws.Connection = conn
	for {
		t, msg, err := ws.Connection.ReadMessage()
		if err != nil {
			break
		}
		err = ws.Connection.WriteMessage(t, msg)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func (ws *Websocket) Broadcast(body string) error {
	if ws.Connection == nil {
		return errors.New("web socket not connected")
	}

	return ws.Connection.WriteJSON(SocketMessage{
		Body:      body,
		CreatedAt: time.Now(),
	})
}
