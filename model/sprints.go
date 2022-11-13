package model

import (
	"time"
)

type Sprint struct {
    SprintID uint `gorm:"primary_key;column:id;type:uint NOT NULL AUTO_INCREMENT" json:"sprintID"`
	ProjectID int `gorm:"not null; foreign_key:ProjectID; references:ProjectID" json:"projectID"`
	SprintName string `gorm:"size:255;not null" json:"sprintName"`
	CreatedAt time.Time
	StartDate time.Time  `gorm:"not null" json:"start_date"`
	EndDate time.Time  `gorm:"not null" json:"end_date"`	
}
