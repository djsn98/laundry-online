package database

import (
	model "LaundryOnlineAPI/model/core"
	"fmt"
)

func Migrate() {
	fmt.Println("Run DB migrate...")
	err := Connection.AutoMigrate(&model.Service{}, &model.Order{}, &model.Customer{})
	if err != nil {
		panic(err)
	}
}
