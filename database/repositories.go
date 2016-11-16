package database

import "gopkg.in/pg.v5"

type (
	UserRepository struct {
		DB *pg.DB
	}

	VCardRepository struct {
		DB *pg.DB
	}
)

func (r *UserRepository) GetByEmail(email string) (User, error) {

	user := new(User)

	err := r.DB.Model(user).Where("LOWER(email) = LOWER(?)", email).Select()

	return user, err
}

func (r *UserRepository) Create(user User) error {
	return r.DB.Insert(user)
}

func (r *VCardRepository) Create(card VCard) error {
	return r.DB.Insert(card)
}
