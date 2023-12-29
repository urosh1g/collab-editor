package repositories

import (
	"github.com/urosh1g/collab-editor/models"
	"gorm.io/gorm"
)

type FileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) Repository[models.File] {
	return &FileRepository{db}
}

func (repo *FileRepository) GetAll() ([]models.File, error) {
	var files []models.File
	result := repo.db.Joins("Project").Find(&files)
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

func (repo *FileRepository) Create(entity any) (models.File, error) {
	file := entity.(models.File)
	result := repo.db.Session(&gorm.Session{FullSaveAssociations: true}).Create(&file)
	if result.Error != nil {
		return models.File{}, result.Error
	}
	return entity.(models.File), nil
}

func (repo *FileRepository) Update(id int64, entity any) (models.File, error) {
	result := repo.db.Save(entity)
	if result.Error != nil {
		return models.File{}, result.Error
	}
	return entity.(models.File), nil
}

func (repo *FileRepository) Delete(id int64) (models.File, error) {
	var file models.File
	result := repo.db.First(&file, id)
	if result.Error != nil {
		return models.File{}, result.Error
	}

	result = repo.db.Delete(&models.File{}, id)
	if result.Error != nil {
		return models.File{}, result.Error
	}
	return file, nil
}
