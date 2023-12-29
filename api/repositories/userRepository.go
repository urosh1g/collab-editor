package repositories

import (
	"github.com/urosh1g/collab-editor/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) Repository[models.User] {
	return &UserRepository{db}
}

func (repo *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	if result := repo.DB.Find(&users); result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (repo *UserRepository) GetOne(id int64) (models.User, error) {
	var user models.User
	if result := repo.DB.First(&user, id); result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func (repo *UserRepository) Create(newUser any) (models.User, error) {
	userRequest := newUser.(models.CreateUserRequest)
	user := models.User{
		Username: userRequest.Username,
		Files:    nil,
		Projects: nil,
	}
	if result := repo.DB.Create(&user); result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func (repo *UserRepository) Update(id int64, entity any) (models.User, error) {
    user := entity.(models.User)
	result := repo.DB.Save(user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func (repo *UserRepository) Delete(id int64) (models.User, error) {
	result, err := repo.GetOne(id)
	if err != nil {
		return models.User{}, err
	}
	if res := repo.DB.Delete(&result, id); res.Error != nil {
		return models.User{}, err
	}
	return result, nil
}
