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
	CreatedAt int64     `gorm:"autoCreateTime" json:"created_at,omitempty"`
	Status    bool      `json:"status"`
	DatasetID uuid.UUID `gorm:"type:uuid" json:"datasetid"`
	Dataset   Dataset   `gorm:"foreignKey:DatasetID"`
}

func addFileToDataSet(c *gin.Context, db *gorm.DB) {
	var file File
	c.BindJSON(&file)
	file.Owner = userid(c, db)
	dbresp := db.Create(&file)
	if dbresp.Error != nil {
		log.Print("Add file to dataset: ", dbresp.Error)
	}
}

func getFilesDataset(c *gin.Context, db *gorm.DB) []File {
	var files []File
	dbid, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		log.Print("Get files dataset: ", err)
	}
	db.Find(&files, File{DatasetID: dbid})
	return files
}
