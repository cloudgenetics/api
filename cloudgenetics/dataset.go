package cloudgenetics

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Dataset DB with dataset
type Dataset struct {
	gorm.Model
	ID        uint64    `gorm:"primaryKey;AUTO_INCREMENT" json:"id, omitempty"`
	UID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name      string    `json: "name, omitentry"`
	Owner     uint
	CreatedAt int64     `gorm:"autoCreateTime"`
	UpdatedAt time.Time `json: "omitentry"`
	Status    bool
	Files     []File `gorm:"many2many:dataset_files;"`
}
