package database

import (
	"gopkg.in/pg.v5"
	"golang.org/x/crypto/bcrypt"
)

type (
	UserRepository struct {
		DB *pg.DB
	}

	VCardRepository struct {
		DB *pg.DB
	}
)

func (r *UserRepository) GetByEmail(email string) (*User, error) {

	user := new(User)

	err := r.DB.Model(user).Where("LOWER(email) = LOWER(?)", email).Select()

	return user, err
}

func (r *UserRepository) Create(user *User) error {

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Update the password with a bcrypt version
	user.Password = string(password)

	return r.DB.Insert(user)
}

func (r *VCardRepository) Create(card VCard) error {
	return r.DB.Insert(card)
}
