package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB map[string]*gorm.DB

func ConnectDatabase() {
	DB = make(map[string]*gorm.DB)

	grammarDsn := "root:123456@tcp(localhost:3306)/Grammar?charset=utf8mb4&parseTime=True&loc=Local"
	comicDsn := "root:123456@tcp(localhost:3306)/Comic?charset=utf8mb4&parseTime=True&loc=Local"
	dbGrammar, err := gorm.Open(mysql.Open(grammarDsn), &gorm.Config{})
	dbComic, err := gorm.Open(mysql.Open(comicDsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	dbGrammar.AutoMigrate(&Blues{}, &Reds{}, &PreExam{}, &Masters{}, &Dictionaries{})
	dbComic.AutoMigrate(&Comics{}, &Contents{})
	DB["grammar"] = dbGrammar
	DB["comic"] = dbComic
}
