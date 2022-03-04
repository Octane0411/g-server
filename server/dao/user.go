package dao

import (
	"g-server/common/db"
	"g-server/server/model"
)

func (d *Dao) CreateUser(username, password, email string) error {
	user := model.User{
		Model:    &model.Model{},
		Username: username,
		Password: password,
		Email:    email,
	}
	return user.Create(db.DB)
}
