package cloudgenetics

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// File DB with datafile info
type File struct {
	gorm.Model
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"omitempty"`
	Name      string    `json: "name"`
	Size      uint      `json: "size"`
	FileType  string    `json: "type"`
	Url       string    `json: "url"`
	Owner     uint      `json:"omitempty"`
	CreatedAt int64     `gorm:"autoCreateTime" json:"omitempty"`
	Status    bool      `json: "status"`
	DatasetID uuid.UUID `gorm: "uniqueIndex;type:uuid" json: "datasetid"`
	Dataset   Dataset   `gorm:"foreignKey:DatasetID"`
}
