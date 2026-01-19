package mysql

import (
	domain "go_2601_04/internal/domain/user"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	db.AutoMigrate(&domain.User{})
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) Update(user *domain.User) error {
	result := r.db.Save(user)
	if result.RowsAffected == 0 {
		return domain.ErrUserNotFound
	}
	return result.Error
}

func (r *UserRepository) Delete(id uint) error {
	result := r.db.Delete(&domain.User{}, id)
	if result.RowsAffected == 0 {
		return domain.ErrUserNotFound
	}
	return result.Error
}

func (r *UserRepository) FindByID(id uint) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, domain.ErrUserNotFound
	}
	return &user, nil
}

func (r *UserRepository) FindAll() ([]domain.User, error) {
	var users []domain.User
	return users, r.db.Find(&users).Error
}
