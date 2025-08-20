package main

import (
	"log"

	"github.com/CrispyChillies/go_authentication/db"

	"github.com/CrispyChillies/go_authentication/models"
)

func main() {
	db.InitDB()
	defer db.DB.Close()
	log.Println("Application started successfully!")
	models.RegisterUser(db.DB, "aaronpham5504", "123456")
}
