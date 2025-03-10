package main

import (
	"go-pos/database"
	_ "go-pos/routers"
	"log"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	// Initialize database connection
	database.Initialize()
	defer database.Close()
	
	// Rest of your application setup
	log.Println("Application started...")
	
	web.Run()
}
