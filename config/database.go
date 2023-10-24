package config

import (
	"golang-crud/helper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DB_USERNAME = "sohatv"
	DB_PASSWORD = "vdKLvIEqtrP4weee"
	DB_NAME     = "test"
	DB_HOST     = "10.5.44.125"
	DB_PORT     = "13306"
)

func DatabaseConnection() *gorm.DB {
	sqlInfo := DB_USERNAME +":"+ DB_PASSWORD +"@tcp"+ "(" + DB_HOST + ":" + DB_PORT +")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)
	return db
}
