package postgres

import (
	"github.com/go-pg/pg/v9"

	"github.com/equimper/meetmeup/models"
)

type UsersRepo struct {
	DB *pg.DB
}

func (u *UsersRepo) GetUserByID(id string) (*models.User, error) {
	var user models.User
	err := u.DB.Model(&user).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}

	return &user, nil
}
