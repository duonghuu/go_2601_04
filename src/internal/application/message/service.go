package message

import "go_2601_04/internal/domain/message"

type Service struct {
	repo message.Repository
}

func NewService(r message.Repository) Service {
	return Service{repo: r}
}

func (s Service) GetHelloMessage(name string) string {
	msg := message.NewMessage("Hello " + name)
	_ = s.repo.Save(msg)
	return msg.Text
}
