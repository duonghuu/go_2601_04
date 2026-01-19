package article

import (
	"errors"
	"strings"
)

type Article struct {
	ID      uint `gorm:"primaryKey"`
	Title   string
	Content string
}

func NewArticle(title, content string) (*Article, error) {
	if strings.TrimSpace(title) == "" {
		return nil, errors.New("title is required")
	}

	return &Article{
		Title:   title,
		Content: content,
	}, nil
}
