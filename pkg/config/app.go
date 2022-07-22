package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)

func Connect(){
	dsn := "admin:Dsci-551@tcp(books.ctvxlirswyn4.us-west-1.rds.amazonaws.com:3306)/dsci551?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", dsn)
	if err != nil{
		panic(err)
	}
	db = d

}

func GetDB() *gorm.DB{
	return db
}

