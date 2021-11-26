package model

import (
	"time"

	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	FolderID  uint
	Folder
}
