package cloudgenetics

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/rds/rdsutils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB postgres GORM object
type DB struct {
	db *gorm.DB
}

// dsn Data Source Name for DB
func dsn() (dsn string, err error) {
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbEndpoint := fmt.Sprintf("%s:%d", dbHost, dbPort)
	region := os.Getenv("AWS_REGION")
	creds := credentials.NewEnvCredentials()
	authToken, err := rdsutils.BuildAuthToken(dbEndpoint, region, dbUser, creds)
	if err != nil {
		panic(err)
	}
	dsn = fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=require password=%s",
		dbHost, dbPort, dbUser, dbName, authToken)
	return dsn, err
}

// DBConnect Connect to Postgres DB
func DBConnect() (db *gorm.DB, err error) {
	dsn, err := dsn()
	if err != nil {
		panic(err)
	}
	db, dberr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dberr != nil {
		panic(dberr)
	}

	// Get generic database object sql.DB to use its functions
	return db, dberr
}
