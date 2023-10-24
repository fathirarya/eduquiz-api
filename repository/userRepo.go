package repository

import (
	"eduquiz-api/model/domain"
	"eduquiz-api/model/schema"
	"eduquiz-api/utils/req"
	"eduquiz-api/utils/res"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *domain.Users) (*domain.Users, error)
	Update(user *domain.Users, id int) (*domain.Users, error)
	FindById(id int) (*domain.Users, error)
	FindByEmail(email string) (*domain.Users, error)
	FindAll() ([]domain.Users, error)
	FindByUsername(name string) (*domain.Users, error)
	Delete(id int) error
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: DB}
}

func (repository *UserRepositoryImpl) Create(user *domain.Users) (*domain.Users, error) {
	userDb := req.UserDomainToUserSchema(*user)
	result := repository.DB.Create(&userDb)
	if result.Error != nil {
		return nil, result.Error
	}

	results := res.UserSchemaToUserDomain(userDb)

	return results, nil
}

func (repository *UserRepositoryImpl) Update(user *domain.Users, id int) (*domain.Users, error) {
	result := repository.DB.Table("users").Where("id = ?", id).Updates(domain.Users{Fullname: user.Fullname, Email: user.Email, Password: user.Password})
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (repository *UserRepositoryImpl) FindById(id int) (*domain.Users, error) {
	user := domain.Users{}

	result := repository.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repository *UserRepositoryImpl) FindByEmail(email string) (*domain.Users, error) {
	user := domain.Users{}

	result := repository.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repository *UserRepositoryImpl) FindAll() ([]domain.Users, error) {
	user := []domain.Users{}

	result := repository.DB.Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (repository *UserRepositoryImpl) FindByUsername(name string) (*domain.Users, error) {
	user := domain.Users{}

	// Menggunakan query LIKE yang tidak case-sensitive
	result := repository.DB.Where("LOWER(name) LIKE LOWER(?)", "%"+name+"%").First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repository *UserRepositoryImpl) Delete(id int) error {
	result := repository.DB.Delete(&schema.Users{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
