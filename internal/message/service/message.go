package service

import (
	"github.com/ahmadkarlam/go-message/internal/message/model"
	"github.com/ahmadkarlam/go-message/internal/message/repository"
	"github.com/ahmadkarlam/go-message/internal/message/repository/memory"
)

type MessageService struct {
	Repository repository.Repository
}

func NewMessageService() MessageService {
	storage := memory.NewMessageStorage()
	return MessageService{
		Repository: memory.NewMessageRepository(storage),
	}
}

func (service *MessageService) GetAllMessage() []model.Message {
	return service.Repository.GetAll()
}

func (service *MessageService) AddMessage(body string) model.Message {
	return service.Repository.Create(body)
}
