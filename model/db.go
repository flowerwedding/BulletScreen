package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var DB *gorm.DB

func init() {
	var err error

	DB, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/dome7?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err.Error())
	}
	if !DB.HasTable(&User{}) {
		if err := DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&User{}).Error; err != nil {
			panic(err)
		}
	}
	if !DB.HasTable(&Bullet{}) {
		if err := DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&Bullet{}).Error; err != nil {
			panic(err)
		}
	}
	if !DB.HasTable(&Word{}) {
		if err := DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&Word{}).Error; err != nil {
			panic(err)
		}
	}
	if !DB.HasTable(&Order{}) {
		if err := DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&Order{}).Error; err != nil {
			panic(err)
		}
	}
}