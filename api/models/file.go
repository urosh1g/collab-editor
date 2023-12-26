package models

import "gorm.io/gorm"

type File struct {
	gorm.Model
	Name string `json:"name"`
	Type string `json:"type"`
	Mode string `json:"mode"`
}
