package main

import "cloudgenetics/cloudgenetics"

func main() {
	r := cloudgenetics.Router()

	cloudgenetics.PublicRoutes(r)
	cloudgenetics.APIV1Routes(r)

	err := r.Run(":5000")
	if err != nil {
		panic(err)
	}
}
