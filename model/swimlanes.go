package model

import (
)

type Swimlane struct {
    SwimlaneID uint `gorm:"primary_key;column:id;type:uint NOT NULL AUTO_INCREMENT" json:"swinlaneID"`
	ProjectID int `gorm:"not null; foreign_key:ProjectID; references:ProjectID" json:"projectID"`
	SwimlaneName string `gorm:"size:255;not null" json:"swimlaneName"`
	SwimlanePositiion int `gorm:"not null" json:"swimlanePosition"`
}
