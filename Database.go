package main

import (
	"github.com/jinzhu/gorm"
	_"github.com/mattn/go-sqlite3"
	"fmt"
)

const (
	dbName = "villasdb.db"
)

func Init(){
	
	db, _ := gorm.Open("sqlite3", dbName)	//openning the database .D
	db.LogMode(true)	//enabling logs in database .
	
	db.DB()
	
	db.DB().Ping()	//set limitations on database connections, etc .
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	//creating the tables .
	//db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Villa{})
	db.CreateTable(&Villa{})
	db.CreateTable(&VillaOwner{})
	db.CreateTable(&Day{})
	db.CreateTable(&Reservation{})
	db.CreateTable(&ReservationTable{})
	db.CreateTable(&ImageGallery{})
	
	fmt.Println("Table Villas has been created .")
}

func dummyVillaData() {
	//creates some dummy villas for testing purposes .
	villaOne := Villa{Name: "Daryakenar 2", Description: "The hottest villa in the north !"}
	villaTwo := Villa{Name: "Kish 1", Description: "The best villa you can find in Kish Island !"}
	villaThree := Villa{Name: "Kish 2", Description: "The luxuriest villa ever in south !"}
	villaFour := Villa{Name: "Antalia 3", Description: "Life is short, engoy it now !!!"}
	
	db, _ := gorm.Open("sqlite3", dbName)	//openning the database .D
	db.LogMode(true)	//enabling logs in database .
	//saving the villas into database:
	db.Save(&villaOne)
	db.Save(&villaTwo)
	db.Save(&villaThree)
	db.Save(&villaFour)	
	
	fmt.Println("Dummy villas has been created .")
}

func getVillas() []Villa {
	db, _ := gorm.Open("sqlite3", dbName)	//openning the database .D
	db.LogMode(true)	//enabling logs in database .
	
	var villas []Villa
	
	db.Select("ID, Name, Description").Find(&villas)
	return villas
}
