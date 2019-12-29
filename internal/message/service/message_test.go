package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/ahmadkarlam/go-message/internal/message/model"
	mocks "github.com/ahmadkarlam/go-message/test/mocks/message"
)

func TestMessageService_GetAllMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mocks.NewMockRepository(ctrl)
	mock.EXPECT().GetAll().Return([]model.Message{
		{Body: "foo"},
		{Body: "bar"},
	}).Times(1)

	service := MessageService{
		Repository: mock,
	}

	messages := service.GetAllMessage()

	assert.Equal(t, 2, len(messages))
	assert.Equal(t, "foo", messages[0].Body)
	assert.Equal(t, "bar", messages[1].Body)
}

func TestMessageService_AddMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mocks.NewMockRepository(ctrl)
	mock.EXPECT().Create("foo").Return(model.Message{Body: "foo"}).Times(1)

	service := MessageService{
		Repository: mock,
	}

	message := service.AddMessage("foo")

	assert.Equal(t, "foo", message.Body)
}
