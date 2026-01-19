package user

type Repository interface {
	Create(user *User) error
	Update(user *User) error
	Delete(id uint) error
	FindByID(id uint) (*User, error)
	FindAll() ([]User, error)
}
