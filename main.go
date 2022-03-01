package main

import (
	db2 "g-server/common/db"
	_ "g-server/common/setting"
	"g-server/server/model"
)

func main() {
	db := db2.DB
	db.Create(&model.User{
		Model:    &model.Model{},
		Uuid:     "",
		Username: "a",
		Password: "b",
		Status:   "",
		Email:    "",
		Avatar:   "",
	})

}
