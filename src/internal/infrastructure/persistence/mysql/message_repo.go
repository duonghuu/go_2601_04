package mysql

import (
	"go_2601_04/internal/domain/message"

	"gorm.io/gorm"
)

// DB Schema (DB Schema)
type Message struct {
	gorm.Model
	Content string
}

func (Message) TableName() string {
    return "messages"
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) message.Repository {
	return messageRepository{db: db}
}

func (r messageRepository) Save(m message.Message) error {
	dbModel := Message{Content: m.Text}
	return r.db.Create(&dbModel).Error
}

func (r messageRepository) GetLast() (*message.Message, error) {
	var result Message
	if err := r.db.Last(&result).Error; err != nil {
		return nil, err
	}
	return &message.Message{Text: result.Content}, nil
}
