package main

import (
	"cloudgenetics/cloudgenetics"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	enverr := godotenv.Load()
	if enverr != nil {
		log.Fatal("Error loading .env file")
	}

	// Create a database
	db, dberr := cloudgenetics.DBConnect()
	if dberr != nil {
		log.Fatal("Failed to connect to database")
		panic(dberr)
	}

	r := cloudgenetics.Router()

	cloudgenetics.PublicRoutes(r)
	cloudgenetics.APIV1Routes(r, db)

	err := r.Run(":5000")
	if err != nil {
		panic(err)
	}
}
