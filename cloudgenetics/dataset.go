package cloudgenetics

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Dataset DB with dataset
type Dataset struct {
	gorm.Model
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"omitempty"`
	Name      string    `json: "name"`
	Owner     uuid.UUID `json: "omitempty"`
	CreatedAt int64     `gorm: "autoCreateTime" json: "omitempty"`
	UpdatedAt time.Time `json: "omitempty"`
	Status    bool      `gorm: "type:boolean;default:true" json: "omitempty"`
	Share     uint      `gorm: "type:uint;default:0" json: "omitempty"`
}

func createDataset(c *gin.Context, db *gorm.DB) uuid.UUID {
	// Get name of dataset
	var ds Dataset
	c.BindJSON(&ds)

	// Fetch userid
	user_id := authid(c)
	// Find user primary key ID
	var user User
	db.Where(&User{Userid: user_id}).First(&user)

	ds.Owner = user.ID
	ds.UpdatedAt = time.Now()
	ds.Status = true

	dbresp := db.Create(&ds)
	if dbresp.Error != nil {
		log.Print("Generate Datset: ", dbresp.Error)
	}
	return ds.ID
}
