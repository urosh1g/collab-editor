package services

import (
	"fmt"

	"github.com/urosh1g/collab-editor/models"
	"github.com/urosh1g/collab-editor/repositories"
)

type ProjectService struct {
	repo repositories.Repository[models.Project]
}

func NewProjectService(repo repositories.Repository[models.Project]) *ProjectService {
	return &ProjectService{repo}
}

func (service *ProjectService) GetAll() ([]models.Project, error) {
	return service.repo.GetAll()
}

func (service *ProjectService) Create(entity models.CreateProjectRequest) (models.Project, error) {
	fmt.Printf("%+v", entity)
	project := models.Project{
		Name: entity.Name,
	}
	return service.repo.Create(project)
}

func (service *ProjectService) Delete(id int64) (models.Project, error) {
	return service.repo.Delete(id)
}
