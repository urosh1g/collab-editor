package models

import "gorm.io/gorm"

type File struct {
	gorm.Model
	Name string
	Type string
	Mode string
}
