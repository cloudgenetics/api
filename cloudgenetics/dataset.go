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
	ID        uint64    `gorm:"primaryKey;AUTO_INCREMENT" json:"id, omitempty"`
	UID       uuid.UUID `gorm:"unique_index"`
	Name      string    `json: "name, omitentry"`
	Owner     uint
	CreatedAt int64     `gorm:"autoCreateTime"`
	UpdatedAt time.Time `json: "omitentry"`
	Status    bool      `gorm: "type:boolean;default:true"`
	Files     []File    `gorm:"many2many:dataset_files;"`
}

type DatasetFile struct {
	FileID    uint64 `gorm:"primaryKey"`
	DatasetID uint64 `gorm:"primaryKey"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type DatasetName struct {
	Name string `json: "datasetname"`
}

func createDataset(c *gin.Context, db *gorm.DB) uuid.UUID {
	var ds Dataset
	// Fetch userid
	user_id := userid(c)
	// Find user primary key ID
	var user User
	db.Where(&User{Userid: user_id}).First(&user)
	// Get name of dataset
	var dsname DatasetName
	c.BindJSON(&dsname)

	datasetid := uuid.New()
	ds.Name = dsname.Name
	ds.UID = datasetid
	ds.Owner = user.ID
	ds.UpdatedAt = time.Now()
	ds.Status = true

	dbresp := db.Create(&ds)
	if dbresp.Error != nil {
		log.Print("Generate Datset: ", dbresp.Error)
	}
	return datasetid
}
