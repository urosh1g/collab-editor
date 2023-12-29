package services

import (
	"errors"

	"github.com/urosh1g/collab-editor/models"
	"github.com/urosh1g/collab-editor/repositories"
)

type UserService struct {
	UserRepository repositories.Repository[models.User]
}

func NewUserService(userRepo repositories.Repository[models.User]) *UserService {
	return &UserService{userRepo}
}

func (service *UserService) GetAll() ([]models.User, error) {
	return service.UserRepository.GetAll()
}

func (service *UserService) GetOne(id int64) (models.User, error) {
	return service.UserRepository.GetOne(id)
}

func (service *UserService) Create(userRequest models.CreateUserRequest) (models.User, error) {
	if userRequest.Username == "" {
		return models.User{}, errors.New("username is required")
	}
	return service.UserRepository.Create(userRequest)
}

func (service *UserService) Delete(id int64) (models.User, error) {
	return service.UserRepository.Delete(id)
}
