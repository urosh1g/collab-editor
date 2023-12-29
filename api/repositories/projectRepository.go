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
	var projects []models.Project
	if result := repo.DB.Find(&projects); result.Error != nil {
		return nil, result.Error
	}
	return projects, nil
}

func (repo *ProjectRepository) GetOne(id int64) (models.Project, error) {
	var project models.Project
	if result := repo.DB.First(&project, id); result.Error != nil {
		return models.Project{}, result.Error
	}
	return project, nil
}

func (repo *ProjectRepository) Create(project any) (models.Project, error) {
	proj := project.(models.Project)
	if result := repo.DB.Create(&proj); result.Error != nil {
		return models.Project{}, result.Error
	}
	return proj, nil
}

func (repo *ProjectRepository) Update(id int64, project any) (models.Project, error) {
	return models.Project{}, nil
}

func (repo *ProjectRepository) Delete(id int64) (models.Project, error) {
	project, err := repo.GetOne(id)
	if err != nil {
		return models.Project{}, err
	}
	if result := repo.DB.Delete(&project, id); result.Error != nil {
		return models.Project{}, err
	}
	return project, nil
}
