package model

import (
	"time"

	"gorm.io/gorm"
)

type Folder struct {
	gorm.Model
	ID              uint
	Name            string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	ParaintFolderID uint
	Children        []Folder `gorm:"foreignkey:ParaintFolderID"`
	Files           []File
}
