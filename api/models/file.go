package models

import "gorm.io/gorm"

type File struct {
	gorm.Model
	Name         string  `json:"name"`
	Type         string  `json:"type"`
	OpenMode     string  `json:"mode"`
	Contributors []*User `json:"contributors,omitempty" gorm:"many2many:user_files;"`
	Project      Project `json:"project" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProjectID    int64   `json:"-"`
}

type CreateFileRequest struct {
	Name      string `json:"name" binding:"required"`
	Type      string `json:"type" binding:"required"`
	OpenMode  string `json:"mode" binding:"required"`
	ProjectID int64  `json:"projectId" binding:"required"`
}
