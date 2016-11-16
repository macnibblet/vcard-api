package database

import (
	"github.com/satori/go.uuid"
	"time"
)

type (
	VCard struct {
		Id       uuid.UUID `sql:"uuid" json:"uuid"`
		UserUuid uuid.UUID `sql:"user_uuid" json:"userUuid"`
		Name     string    `sql:"name" json:"name"`
		Path     string    `sql:"path" json:"path"`
	}

	User struct {
		Id        uuid.UUID `sql:"uuid" json:"uuid"`
		Email     string    `sql:"email" json:"email"`
		Password  string    `sql:"password" json:"-"`
		CreatedAt time.Time `sql:"created_at" json:"createdAt"`
		UpdatedAt time.Time `sql:"updated_at" json:"updatedAt"`
	}
)
