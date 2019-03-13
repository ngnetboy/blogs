package service

import (
	"log"
	"model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func ConnectDB() {
	var err error
	db, err = gorm.Open("sqlite3", model.Conf.DB)
	// db, err = gorm.Open("sqlite3", "blogs.db")
	if err != nil {
		log.Fatalln("Open database failed:", err.Error())
	}

	if err = db.AutoMigrate(model.Models...).Error; nil != err {
		log.Fatal("auto migrate tables failed: " + err.Error())
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(50)

}

func DisconnectDB() {
	if db != nil {
		if err := db.Close(); nil != err {
			log.Fatalln("Disconnect from database failed:", err.Error())
		}
	}
}
