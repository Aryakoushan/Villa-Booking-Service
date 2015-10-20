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
	ID 			int
	Owner 		VillaOwner
	Name 		string `sql:"size:255"` //Default size for string .
	Description string `sql:"size:255"`
	Table		ReservationTable
	Gallery		ImageGallery
	
	CreatedAt    time.Time
    UpdatedAt    time.Time
    DeletedAt    *time.Time
}

type ReservationTable struct {
	gorm.Model
	ID				int
	Reservations	[]Reservation
	
	CreatedAt    time.Time
    UpdatedAt    time.Time
    DeletedAt    *time.Time
}

type ImageGallery struct {
	gorm.Model
	ID			int
	PicUrls		[]string
	
	CreatedAt    time.Time
    UpdatedAt    time.Time
    DeletedAt    *time.Time
}

type VillaOwner struct {
	gorm.Model
	ID			int
	Name		string	`sql:"size:255"`
	Villas		[]Villa
	Pic			string	`sql:"size:255"`
	Note		string	`sql:"size:255"`
	
	CreatedAt    time.Time
    UpdatedAt    time.Time
    DeletedAt    *time.Time
}

type Day struct {
	gorm.Model
	ID			int
	ColorStat	string
	price		float64
	Date		time.Time
	VillaID		int
}

type Reservation struct {
	gorm.Model
	ID			int
	AgentID		int
	CustomerID	int
	Days		[]Day
	TotalPrice	float64
	PaymentStat	string
	
	CreatedAt    time.Time
    UpdatedAt    time.Time
    DeletedAt    *time.Time
}