package database

import model "LaundryOnlineAPI/model/core"

func Migrate() {
	err := Connection.AutoMigrate(&model.Service{}, &model.Order{}, &model.Customer{})
	if err != nil {
		panic(err)
	}
}
