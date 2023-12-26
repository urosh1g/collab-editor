package repositories

import (
	"gorm.io/gorm"
	"github.com/urosh1g/collab-editor/models"
)

type FileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) FileRepository {
	return FileRepository{db}
}

func (repo *FileRepository) GetAll() ([]models.File, error) {
	var files []models.File
	result := repo.db.Find(&files)
	if result.Error != nil {
		return nil, result.Error
	}
	return files, nil
}

func (repo *FileRepository) GetOne(id int64) (models.File, error) {
	var file models.File
	result := repo.db.First(&file, id)
	if result.Error != nil {
		return models.File{}, result.Error
	}
	return file, nil
}
