package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DbConnection() *gorm.DB {
	dsn := "root:recce2021@tcp(127.0.0.1:3306)/payment?"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("connected to db")
	return db

}
