package main

import (
	"gorm.io/gorm"
	"fmt"
	"gorm.io/driver/postgres"
)

func GetDatabase(config *DbConfig) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.GetDSN()), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("GetDatabase failure: %w", err))
	}
	return db
}
