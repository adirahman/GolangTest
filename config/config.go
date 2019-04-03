package config

import (
	"payfazztest/structs"

	"github.com/jinzhu/gorm"
)

// DBInit create connection to database
func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/go_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(structs.Person{})
	db.AutoMigrate(structs.Company{})
	db.AutoMigrate(structs.Friend{})
	db.AutoMigrate(structs.WorkHistory{})
	return db
}