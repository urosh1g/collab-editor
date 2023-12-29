package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name  string  `json:"name"`
	Files []File  `json:"files,omitempty"`
	Users []*User `json:"users,omitempty" gorm:"many2many:project_users"`
}

type CreateProjectRequest struct {
	Name string `json:"name"`
}
