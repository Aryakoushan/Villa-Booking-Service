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
<<<<<<< HEAD
=======
	
>>>>>>> 0b7e19c6bd11d2abcd946782ded75bc7a8cdd95d
	Owner 		VillaOwner
	Name 		string `sql:"size:255"` //Default size for string .
	Description string `sql:"size:255"`
	Table		ReservationTable
	Gallery		ImageGallery
}

type ReservationTable struct {
	gorm.Model
<<<<<<< HEAD
=======
	
>>>>>>> 0b7e19c6bd11d2abcd946782ded75bc7a8cdd95d
	Reservations	[]Reservation
}

type ImageGallery struct {
	gorm.Model
<<<<<<< HEAD
=======
	
>>>>>>> 0b7e19c6bd11d2abcd946782ded75bc7a8cdd95d
	PicUrls		[]string
}

type VillaOwner struct {
	gorm.Model
<<<<<<< HEAD
=======
	
>>>>>>> 0b7e19c6bd11d2abcd946782ded75bc7a8cdd95d
	Name		string	`sql:"size:255"`
	Villas		[]Villa
	Pic			string	`sql:"size:255"`
	Note		string	`sql:"size:255"`
}

type Day struct {
	gorm.Model
<<<<<<< HEAD
=======
	
>>>>>>> 0b7e19c6bd11d2abcd946782ded75bc7a8cdd95d
	ColorStat	string
	price		float64
	Date		time.Time
	VillaID		int
}

type Reservation struct {
	gorm.Model
<<<<<<< HEAD
=======
	
>>>>>>> 0b7e19c6bd11d2abcd946782ded75bc7a8cdd95d
	AgentID		int
	CustomerID	int
	Days		[]Day
	TotalPrice	float64
	PaymentStat	string
}