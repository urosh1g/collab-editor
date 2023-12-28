package services

import (
	"github.com/urosh1g/collab-editor/models"
	"github.com/urosh1g/collab-editor/repositories"
)

type UserService struct {
	userRepo    repositories.Repository[models.User]
	projectRepo repositories.Repository[models.Project]
	fileRepo	repositories.Repository[models.File]
}

func NewUserService(UserRepo repositories.Repository[models.User], projectRepo repositories.Repository[models.Project]) *UserService {
	return &UserService{UserRepo, projectRepo}
}

func (service *UserService) GetAll() ([]models.User, error) {
	return service.UserRepo.GetAll()
}

func (service *UserService) GetOne(id int64) (models.User, error) {
	return service.UserRepo.GetOne(id)
}

func (service *UserService) Create(User *models.User) (models.User, error) {
	result, err := service.projectRepo.GetOne(User.ProjectID)
	if err != nil {
		return models.User{}, err
	}
	service.UserRepo.(*repositories.UserRepository).DB.Association("Projects").Append(&result)
	service.UserRepo.(*repositories.UserRepository).DB.Association("FileContributions").Append(&result)
	_ = result
	return service.UserRepo.Create(User)
}