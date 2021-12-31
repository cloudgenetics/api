package cloudgenetics

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/rds/rdsutils"
	_ "github.com/lib/pq"

	"github.com/jinzhu/gorm"
	// Importing postgres dialects
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func db() {
	dbName := "csgdevdb"
	dbUser := "dev"
	dbHost := "csgdevdb.ctzmtn99czvb.us-east-1.rds.amazonaws.com"
	dbPort := 5432
	dbEndpoint := fmt.Sprintf("%s:%d", dbHost, dbPort)
	region := "us-east-1"
	creds := credentials.NewEnvCredentials()
	authToken, err := rdsutils.BuildAuthToken(dbEndpoint, region, dbUser, creds)
	if err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=require password=%s",
		dbHost, dbPort, dbUser, dbName, authToken)

	db, dberr := gorm.Open("postgres", dsn)
	if dberr != nil {
		panic(err)
	}
	// Get generic database object sql.DB to use its functions
	err = db.DB().Ping()
	if err != nil {
		// db.Close()
		panic(err)
	} else {
		fmt.Println("Connected")
	}
	db.Close()
}
