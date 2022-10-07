package config


import (
	"github.com/jinzhu/gorm"
	// "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "pixelle:areapassword@tcp(localhost:3306)/AREA?parseTime=true")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDb() * gorm.DB {
	return db
}