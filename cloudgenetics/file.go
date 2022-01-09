package cloudgenetics

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// File DB with datafile info
type File struct {
	gorm.Model
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"omitempty"`
	Name      string    `json:"name"`
	Size      uint      `json:"size"`
	FileType  string    `json:"type"`
	Url       string    `json:"url"`
	Owner     uuid.UUID `json:"omitempty"`
	CreatedAt int64     `gorm:"autoCreateTime" json:"omitempty"`
	Status    bool      `json:"status"`
	DatasetID uuid.UUID `gorm:"uniqueIndex;type:uuid" json:"datasetid"`
	Dataset   Dataset   `gorm:"foreignKey:DatasetID"`
}

func addFileToDataSet(c *gin.Context, db *gorm.DB) {
	var file File
	c.BindJSON(&file)
	dbresp := db.Create(&file)
	if dbresp.Error != nil {
		log.Print("Add file: ", dbresp.Error)
	}
}
