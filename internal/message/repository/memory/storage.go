package memory

import (
	"github.com/ahmadkarlam/go-message/internal/message/model"
)

var (
	Storage *MessageStorage
)

type MessageStorage struct {
	Messages []model.Message
}

func NewMessageStorage() *MessageStorage {
	if Storage == nil {
		Storage = &MessageStorage{Messages: []model.Message{}}
	}
	return Storage
}
