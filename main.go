package main

import (
	"arqsoft_proyecto/app"
	"arqsoft_proyecto/db"
)

func main() {
	// Initialize the database connection
	database := db.InitConnection()
	app.StartRoute()

	// Perform any necessary operations with the database
	// For example, you can create a new user or perform queries

	// dbClose the database connection when done
	defer db.Close(database)
}
