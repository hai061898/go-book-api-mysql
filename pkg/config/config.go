package config

import(
		"github.com/jinzhu/gorm"
		"github.com/jinzhu/gorm/dialects/mysql"
)


var (
	db * gorm.DB
)

func Connect(){
	d, err := gorm.Open("mysql","")
	if err != nil {
		panic(err)
	}
	db = d 
}

func GetDB() *gorm.DB{
	return db
}


