package message

type Repository interface {
	Save(msg Message) error
	GetLast() (*Message, error)
}
