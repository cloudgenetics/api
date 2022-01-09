package cloudgenetics

import (
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User DB User table
type User struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"omitempty"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Userid      string    `gorm:"unique_index" json:"userid"`
	EmailVerify bool      `json:"email_verify"`
	UpdatedAt   time.Time `json:"updated_at"`
	Role        uint      `json:"role, omitentry"`
	Active      bool      `gorm:"default:true"`
}

// registerUser Register a new user in DB
func registerUser(c *gin.Context, db *gorm.DB) string {
	// Create a new user
	var user User
	if err := c.BindJSON(&user); err != nil {
		return "JSON bind failed"
	}

	// Fetch userid from authentication
	userid := authid(c)
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

// authid returns userid from JWT auth token
func authid(c *gin.Context) string {
	// Get token from http request, parse JWT to get "sub" (user id)
	return c.Request.Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["sub"].(string)
}

func userid(c *gin.Context, db *gorm.DB) uuid.UUID {
	// Fetch userid
	user_id := authid(c)
	// Find user primary key ID
	var user User
	db.Where(&User{Userid: user_id}).First(&user)
	return user.ID
}
