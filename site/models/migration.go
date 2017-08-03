package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"path/to/your/app/config"
)

var DB, _ = gorm.Open("mysql", config.GetS("dbconfig"))

func init() {
	DB.SingularTable(true)
	DB.AutoMigrate(&User{})
}