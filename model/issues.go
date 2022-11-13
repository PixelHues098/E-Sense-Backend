package model

import (
	"time"
)

type Issue struct {
    Issues uint `gorm:"primary_key;column:id;type:uint NOT NULL AUTO_INCREMENT" json:"issues"`
	ProjectID int `gorm:"not null; foreign_key:ProjectID; references:ProjectID" json:"projectID"`
	ReporterID int `gorm:"not null; foreign_key:ReporterID; references:ID" json:"reporterID"`
	AssigneeID int `gorm:"not null; foreign_key:AssigneeID; references:ID" json:"assigneeID"`
	SprintID int `gorm:"not null; foreign_key:SprintID; references:SprintID" json:"sprintID"`
	CreatedAt time.Time	
	IssueType int `gorm:"not null" json:"issueType"`
	PriorityType int `gorm:"not null" json:"priorityType"`
}
