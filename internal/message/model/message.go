package model

type Message struct {
	Body string `json:"body"`
}

func NewMessage(body string) Message {
	return Message{
		Body: body,
	}
}
