package mysql

import (
	domain "go_2601_04/internal/domain/article"

	"gorm.io/gorm"
)

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

func (r *ArticleRepository) Create(article *domain.Article) error {
	return r.db.Create(article).Error
}

func (r *ArticleRepository) Update(article *domain.Article) error {
	result := r.db.Save(article)
	if result.RowsAffected == 0 {
		return domain.ErrArticleNotFound
	}
	return result.Error
}

func (r *ArticleRepository) Delete(id uint) error {
	result := r.db.Delete(&domain.Article{}, id)
	if result.RowsAffected == 0 {
		return domain.ErrArticleNotFound
	}
	return result.Error
}

func (r *ArticleRepository) FindByID(id uint) (*domain.Article, error) {
	var article domain.Article
	if err := r.db.First(&article, id).Error; err != nil {
		return nil, domain.ErrArticleNotFound
	}
	return &article, nil
}

func (r *ArticleRepository) FindAll() ([]domain.Article, error) {
	var articles []domain.Article
	return articles, r.db.Find(&articles).Error
}
