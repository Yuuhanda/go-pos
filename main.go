package main

import (
	"go-pos/database"
	"log"
)

func main() {
	// Initialize database connection
	database.Initialize()
	defer database.Close()
	
	// Rest of your application setup
	log.Println("Application started...")
	
	// Your application code here...
}
