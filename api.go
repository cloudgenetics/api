package main

import "cloudgenetics/cloudgenetics"

func main() {
	r := cloudgenetics.Router()

	cloudgenetics.PublicRoutes(r)
	cloudgenetics.APIV1Routes(r)

	err := r.Run(":4000")
	if err != nil {
		panic(err)
	}
}
