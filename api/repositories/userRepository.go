package repositories

import (
	"github.com/urosh1g/collab-editor/models"
	"gorm.io/gorm"
)

type UserRepository struct{
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) *UserRepository{
	return &UserRepository{DB}
}

func (repo *UserRepository) GetAll() ([]models.User, error) {
	var Users []models.User
	result := repo.DB.Find(&Users)
	if result.Error != nil {
		return nil, result.Error
	}
	return Users, nil
}

func (repo *UserRepository) GetOne(id int64) (models.User, error) {
	var User models.User
	result := repo.DB.First(&User, id)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return User, nil
}

func (repo *UserRepository) Create(entity *models.User) (models.User, error) {
	result := repo.DB.Omit("Projects").Omit("FileContirubions").Create(&entity)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return *entity, nil
}

func (repo *UserRepository) Update(id int64, entity *models.User) (models.User, error) {
	result := repo.DB.Save(entity)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return *entity, nil
}

func (repo *UserRepository) Delete(id int64) (models.User, error) {
	var User models.User
	result := repo.DB.First(&User, id)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	result = repo.DB.Delete(&models.User{}, id)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return User, nil
}