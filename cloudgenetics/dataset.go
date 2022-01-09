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
	Name      string    `json:"name"`
	CreatedAt int64     `gorm:"autoCreateTime" json:"omitempty"`
	UpdatedAt time.Time `json:"omitempty"`
	Status    bool      `gorm:"type:boolean;default:true" json:"omitempty"`
	Share     uint      `gorm:"type:uint;default:0" json:"omitempty"`
	OwnerID   uuid.UUID `gorm:"uniqueIndex;type:uint" json:"omitempty"`
	User      User      `gorm:"foreignKey:OwnerID"`
}

func createDataset(c *gin.Context, db *gorm.DB) uuid.UUID {
	// Get name of dataset
	var ds Dataset
	c.BindJSON(&ds)

	ds.OwnerID = userid(c, db)
	ds.UpdatedAt = time.Now()
	ds.Status = true

	dbresp := db.Save(&ds)
	if dbresp.Error != nil {
		log.Print("Generate Datset: ", dbresp.Error)
	}
	return ds.ID
}
