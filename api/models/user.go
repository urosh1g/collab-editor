package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username          string    `json:"username"`
	Projects          []Project `json:"projects" gorm:"many2many:project_users;"`
	FileContributions []File    `json:"fileContributions" gorm:"many2many:file_contributors"`
}