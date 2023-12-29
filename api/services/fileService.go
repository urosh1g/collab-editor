package services

import (
	"github.com/urosh1g/collab-editor/models"
	"github.com/urosh1g/collab-editor/repositories"
)

type FileService struct {
	fileRepo    repositories.Repository[models.File]
	projectRepo repositories.Repository[models.Project]
}

func NewFileService(fileRepo repositories.Repository[models.File], projectRepo repositories.Repository[models.Project]) *FileService {
	return &FileService{fileRepo, projectRepo}
}

func (service *FileService) GetAll() ([]models.File, error) {
	return service.fileRepo.GetAll()
}

func (service *FileService) Create(userRequest models.CreateFileRequest) (models.File, error) {
	project, err := service.projectRepo.GetOne(userRequest.ProjectID)
	if err != nil {
		return models.File{}, err
	}
	file := models.File{
		Name:         userRequest.Name,
		Type:         userRequest.Type,
		OpenMode:     userRequest.OpenMode,
		Contributors: nil,
		Project:      project,
	}
	_ = file
	return service.fileRepo.Create(file)
}

func (service *FileService) Delete(id int64) (models.File, error) {
	return service.fileRepo.Delete(id)
}
