package main

import (
	"time"
	"github.com/jinzhu/gorm"
)

type Model struct {
	ID			int `gorm:"primary_key"`
	CreatedAt  	time.Time
    UpdatedAt  	time.Time
    DeletedAt  	*time.Time
}

type Villa struct {
	gorm.Model
	
	Owner 		VillaOwner
	Name 		string `sql:"size:255"` //Default size for string .
	Description string `sql:"size:255"`
	Table		ReservationTable
	Gallery		ImageGallery
	Rating		float32
	Price		float64
}

type ReservationTable struct {
	gorm.Model
	
	Reservations	[]Reservation
}

type ImageGallery struct {
	gorm.Model
	
	PicUrls		[]string
}

type VillaOwner struct {
	gorm.Model
	
	Name		string	`sql:"size:255"`
	Villas		[]Villa
	Pic			string	`sql:"size:255"`
	Note		string	`sql:"size:255"`
}

type Day struct {
	gorm.Model
	
	ColorStat	string
	price		float64
	Date		time.Time
	VillaID		int
}

type Reservation struct {
	gorm.Model
	
	AgentID		int
	CustomerID	int
	Days		[]Day
	TotalPrice	float64
	PaymentStat	string
}