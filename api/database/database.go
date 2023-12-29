package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urosh1g/collab-editor/configuration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func init() {
	config := configuration.GetConfig()
	Db = GetDatabase(&config)
}

func GetDatabase(config *configuration.DbConfig) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)
	db, err := gorm.Open(postgres.Open(config.GetDSN()), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(fmt.Errorf("GetDatabase failure: %w", err))
	}
	return db
}
