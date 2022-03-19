package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DbConnection() *gorm.DB {
	dsn := "root:password@tcp(127.0.0.1:3306)/books?"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("connected to db")
	return db

}
