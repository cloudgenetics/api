package cloudgenetics

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/rds/rdsutils"
	_ "github.com/lib/pq"

	"github.com/jinzhu/gorm"
	// Importing postgres dialects
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB postgres GORM object
type DB struct {
	db *gorm.DB
}

// dbConfig Configure the DB connection settings
func dbConfig() (dsn string, err error) {
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

// dbConnect Connect to Postgres DB
func Connect() (db *gorm.DB, err error) {
	dsn, err := dbConfig()
	if err != nil {
		panic(err)
	}
	db, dberr := gorm.Open("postgres", dsn)
	if dberr != nil {
		panic(dberr)
	}

	// Get generic database object sql.DB to use its functions
	pingerror := db.DB().Ping()
	if pingerror != nil {
		db.Close()
		panic(pingerror)
	} else {
		fmt.Println("Connected")
	}
	return db, dberr
}
