package article

type Repository interface {
	Create(user *Article) error
	Update(user *Article) error
	Delete(id uint) error
	FindByID(id uint) (*Article, error)
	FindAll() ([]Article, error)
}
