package article

import domain "go_2601_04/internal/domain/article"

type Service struct {
	repo domain.Repository
}

func NewArticleService(repo domain.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(name, email string) (*domain.Article, error) {
	article, err := domain.NewArticle(name, email)
	if err != nil {
		return nil, err
	}
	return article, s.repo.Create(article)
}

func (s *Service) Update(id uint, name, email string) (*domain.Article, error) {
	article, err := domain.NewArticle(name, email)
	if err != nil {
		return nil, err
	}
	article.ID = id
	return article, s.repo.Update(article)
}

func (s *Service) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *Service) Get(id uint) (*domain.Article, error) {
	return s.repo.FindByID(id)
}

func (s *Service) List() ([]domain.Article, error) {
	return s.repo.FindAll()
}
