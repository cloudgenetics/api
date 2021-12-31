package cloudgenetics

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// User DB User table
type User struct {
	ID          uint      `gorm:"primaryKey;AUTO_INCREMENT" json:"id, omitempty"`
	Name        string    `json: "name"`
	Email       string    `json: "email"`
	Userid      string    `gorm:"unique_index" json: "userid"`
	EmailVerify bool      `json: "email_verify"`
	UpdatedAt   time.Time `json: "updated_at"`
	Role        uint      `json: "role, omitentry"`
}

// registerUser Register a new user in DB
func registerUser(c *gin.Context, db *gorm.DB) string {
	// Create a new user
	var user User

	// Fetch userid from authentication
	userid := userid(c)
	if err := c.BindJSON(&user); err != nil {
		return "JSON bind failed"
	}
	// Check if the authenticated userid
	if user.Userid == userid {
		// Check if user already exists
		var users []User
		db.Where(&User{Userid: userid}).Find(&users)
		if len(users) == 0 {
			// Create user in db
			result := db.Create(&user)
			if result.Error != nil {
				return "create DB failed"
			} else {
				return "success"
			}
		} else {
			return "User already exists"
		}

	} else {
		return "Invalid userid, id doesn't match"
	}
}
