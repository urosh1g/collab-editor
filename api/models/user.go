package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string     `json:"username"`
	Projects []*Project `json:"projects,omitempty" gorm:"many2many:project_users"`
	Files    []*File    `json:"files,omitempty" gorm:"many2many:user_files"`
}
