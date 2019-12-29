package memory

import (
	"github.com/ahmadkarlam/go-message/internal/message/model"
)

type MessageStorage struct {
	Messages []model.Message
}

func NewMessageStorage() *MessageStorage {
	return &MessageStorage{Messages: []model.Message{}}
}
