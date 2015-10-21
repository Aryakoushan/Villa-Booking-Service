// villa_service_3 project main.go
package main

import (
	"fmt"
	"bytes"
	"github.com/go-martini/martini"
	//"github.com/martini-contrib/auth"
	"encoding/json"
)

func main() {
	fmt.Println("Villa Booking API starting ...")
	//Init()	//creates the database and the tables .
	//dummyVillaData()	//creates some dummty villas for testing purposes .
	m := martini.Classic()
	//m.Use(auth.Basic("root", "secret"))
	
	m.Get("/", Index)
	m.Get("/getAllVillas", GetAllVillas)
	m.Run()
}

func Index() string{
	return "Welcome to Villa Booking API !"
}

func GetAllVillas() string {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(getVillas())
	return buf.String()
}
