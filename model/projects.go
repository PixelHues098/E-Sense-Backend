package model

import (
	"time"
)

type Project struct {
    ProjectID uint `gorm:"primary_key;column:id;type:uint NOT NULL AUTO_INCREMENT" json:"projectID"`
	CreatedAt time.Time
	OwnerID string `gorm:"size:255;not null" json:"ownerID"`
	
}
