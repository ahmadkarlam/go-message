package memory

import (
	"github.com/ahmadkarlam/go-message/internal/message/model"
	"github.com/ahmadkarlam/go-message/internal/message/repository"
)

type MessageRepository struct {
	storage MessageStorage
}

func NewMessageRepository(storage MessageStorage) repository.Repository {
	return &MessageRepository{storage: storage}
}

func (repository *MessageRepository) GetAll() []model.Message {
	return repository.storage.Messages
}

func (repository *MessageRepository) Create(body string) model.Message {
	message := model.NewMessage(body)
	repository.storage.Messages = append(repository.storage.Messages, message)

	return message
}
