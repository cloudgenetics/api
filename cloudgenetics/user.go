package cloudgenetics

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// User DB User table
type User struct {
	ID          uint      `gorm:"primaryKey;AUTO_INCREMENT" json:"id, omitempty"`
	Name        string    `json: "name"`
	Email       string    `json: "email"`
	Userid      string    `gorm:"unique_index" json: "userid"`
	EmailVerify bool      `json: "emailverify, omitentry"`
	CreatedAt   time.Time `json: "createdat, omitentry"`
	UpdatedAt   int       `json: "updatedat, omitentry"`
}

// registerUser Register a new user in DB
func registerUser(c *gin.Context) {
	var user User
	userid := userid(c)
	if err := c.BindJSON(&user); err != nil {
		log.Print("Bind user: ", user, err)
	}
	if user.Userid != userid {
		log.Print("Invalid user id")
	}
	log.Print("User: ", user)
}
