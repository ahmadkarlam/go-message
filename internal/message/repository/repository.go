package repository

import "github.com/ahmadkarlam/go-message/internal/message/model"

type Repository interface {
	GetAll() []model.Message
	Create(body string) model.Message
}
