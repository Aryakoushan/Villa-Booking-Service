// villa_service_3 project sql helper functions
package main

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB

func initialize_variables(name string, variable * string){
	fmt.Print("Enter the " + name + ". Default is ", *variable, ": ")
	var temp string
	fmt.Scanf("%s", &temp)
	if temp != "" {
		*variable = temp
	}
}

func init(){
	
	initialize_variables("host",     &DBHost)
	initialize_variables("port",     &DBPort)
	initialize_variables("username", &DBUser)
	initialize_variables("password", &DBPass)
	initialize_variables("dbname",   &DBName)
	
	database := "user=" + DBUser + " password=" + DBPass + " dbname=" + DBName 
	database += " host=" + DBHost + " port=" + DBPort + " sslmode=disable"

	db, err := sql.Open("postgres", database)
	
	err = db.Ping()
	if err != nil {
		log.Println("Can not open the database. Check you inputs and run program again.")
		log.Fatal(err)
	}else{
		log.Println("Database is ready")
	}
	
	
}