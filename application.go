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

	r := cloudgenetics.Router()

	cloudgenetics.PublicRoutes(r)
	cloudgenetics.APIV1Routes(r)

	err := r.Run(":5000")
	if err != nil {
		panic(err)
	}
}
