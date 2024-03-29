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
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"uuid,omitempty"`
	Name      string    `json:"name"`
	CreatedAt int64     `gorm:"autoCreateTime" json:"omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Status    bool      `gorm:"type:boolean;default:true" json:"status,omitempty"`
	Share     uint      `gorm:"type:uint;default:0" json:"share,omitempty"`
	OwnerID   uuid.UUID `gorm:"type:uint" json:"omitempty"`
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
		log.Print("Generate Dataset: ", dbresp.Error)
	}
	return ds.ID
}

func addDataset(ds Dataset, db *gorm.DB) uuid.UUID {
	dbresp := db.Save(&ds)
	if dbresp.Error != nil {
		log.Print("Add Dataset: ", dbresp.Error)
	}
	return ds.ID
}

func listDatasets(c *gin.Context, db *gorm.DB) []Dataset {
	var datasets []Dataset
	userid := userid(c, db)
	db.Find(&datasets, Dataset{OwnerID: userid})
	return datasets
}

func getDataset(c *gin.Context, db *gorm.DB) Dataset {
	var ds Dataset
	dbid, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		log.Print("Get dataset: ", err)
	}
	db.First(&ds, Dataset{ID: dbid})
	return ds
}
