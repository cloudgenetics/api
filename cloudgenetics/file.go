package cloudgenetics

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// File DB with datafile info
type File struct {
	gorm.Model
	ID        uint64    `gorm:"primaryKey;AUTO_INCREMENT" json:"id, omitempty"`
	UID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name      string    `json: "name"`
	Size      uint      `json: "size"`
	FileType  string
	Loc       string
	Owner     uint
	CreatedAt int64 `gorm:"autoCreateTime"`
	Status    bool
}
