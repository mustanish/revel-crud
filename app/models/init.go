package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/revel/revel"
)

func init() {
	revel.OnAppStart(InitDB)
}

var DB *gorm.DB

func InitDB() {
	info, found := revel.Config.String("db.info")
	if !found {
		revel.AppLog.Fatal("Could not find dbInfo")
	}
	db, err := gorm.Open("mysql", info)
	if err != nil {
		revel.AppLog.Fatal("Could not connect to database with error: " + err.Error())
	}
	db.DB()
	db.AutoMigrate(&User{})
	DB = db
}
