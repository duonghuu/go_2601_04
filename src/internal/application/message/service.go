package message

import "go_2601_04/internal/domain/message"

type Service struct{}

func NewService() Service {
	return Service{}
}

func (s Service) GetHelloMessage(name string) string {
	msg := message.NewMessage("Hello " + name)
	return msg.Text
}
