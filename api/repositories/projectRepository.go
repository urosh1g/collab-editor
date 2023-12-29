package repositories

import (
	"github.com/urosh1g/collab-editor/models"
	"gorm.io/gorm"
)

type ProjectRepository struct {
	DB *gorm.DB
}

func NewProjectRepository(db *gorm.DB) Repository[models.Project] {
	return &ProjectRepository{db}
}

func (repo *ProjectRepository) GetAll() ([]models.Project, error) {
	return nil, nil
}

func (repo *ProjectRepository) GetOne(id int64) (models.Project, error) {
	return models.Project{}, nil
}

func (repo *ProjectRepository) Create(project any) (models.Project, error) {
	return models.Project{}, nil
}

func (repo *ProjectRepository) Update(id int64, project any) (models.Project, error) {
	return models.Project{}, nil
}

func (repo *ProjectRepository) Delete(id int64) (models.Project, error) {
	return models.Project{}, nil
}
