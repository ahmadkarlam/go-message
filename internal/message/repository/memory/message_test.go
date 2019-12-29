package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ahmadkarlam/go-message/internal/message/model"
)

func TestMessageRepository_GetAll(t *testing.T) {
	storage := NewMessageStorage()
	storage.Messages = []model.Message{
		{Body: "foo"},
		{Body: "bar"},
	}

	repository := NewMessageRepository(storage)

	messages := repository.GetAll()

	assert.Equal(t, 2, len(messages))
	assert.Equal(t, "foo", messages[0].Body)
	assert.Equal(t, "bar", messages[1].Body)
}

func TestMessageRepository_Create(t *testing.T) {
	storage := NewMessageStorage()

	repository := NewMessageRepository(storage)

	_ = repository.Create("foo")
	messages := storage.Messages

	assert.Equal(t, 1, len(messages))
	assert.Equal(t, "foo", messages[0].Body)
}
