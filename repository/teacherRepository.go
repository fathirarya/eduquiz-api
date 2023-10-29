package repository

import (
	"eduquiz-api/model/domain"
	"eduquiz-api/utils/req"
	"eduquiz-api/utils/res"

	"gorm.io/gorm"
)

type TeacherRepository interface {
	CreateTeacher(teacher *domain.Teacher) (*domain.Teacher, error)
	UpdateTeacher(teacher *domain.Teacher, id int) (*domain.Teacher, error)
	FindById(id int) (*domain.Teacher, error)
	FindByEmail(email string) (*domain.Teacher, error)
	FindAll() ([]domain.Teacher, error)
	FindByUsername(name string) (*domain.Teacher, error)
	Delete(id int) error
}

type TeacherRepositoryImpl struct {
	DB *gorm.DB
}

func NewTeacherRepository(DB *gorm.DB) TeacherRepository {
	return &TeacherRepositoryImpl{DB: DB}
}

func (repository *TeacherRepositoryImpl) CreateTeacher(teacher *domain.Teacher) (*domain.Teacher, error) {
	teacherDb := req.TeacherDomainToTeacherSchema(*teacher)
	result := repository.DB.Create(&teacherDb)
	if result.Error != nil {
		return nil, result.Error
	}

	results := res.TeacherSchemaToTeacherDomain(teacherDb)

	return results, nil
}

func (repository *TeacherRepositoryImpl) UpdateTeacher(teacher *domain.Teacher, id int) (*domain.Teacher, error) {
	result := repository.DB.Table("teachers").Where("id = ?", id).Updates(domain.Teacher{Fullname: teacher.Fullname, Email: teacher.Email, Password: teacher.Password})
	if result.Error != nil {
		return nil, result.Error
	}

	return teacher, nil
}

func (repository *TeacherRepositoryImpl) FindById(id int) (*domain.Teacher, error) {
	teacher := domain.Teacher{}

	result := repository.DB.First(&teacher, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &teacher, nil
}

func (repository *TeacherRepositoryImpl) FindByEmail(email string) (*domain.Teacher, error) {
	teacher := domain.Teacher{}

	result := repository.DB.Where("email = ?", email).First(&teacher)
	if result.Error != nil {
		return nil, result.Error
	}

	return &teacher, nil
}

func (repository *TeacherRepositoryImpl) FindAll() ([]domain.Teacher, error) {
	var teachers []domain.Teacher

	result := repository.DB.Find(&teachers)
	if result.Error != nil {
		return nil, result.Error
	}

	return teachers, nil
}

func (repository *TeacherRepositoryImpl) FindByUsername(username string) (*domain.Teacher, error) {
	teacher := domain.Teacher{}

	result := repository.DB.Where("LOWER(username) LIKE LOWER(?)", "%"+username+"%").First(&teacher)
	if result.Error != nil {
		return nil, result.Error
	}

	return &teacher, nil
}

func (repository *TeacherRepositoryImpl) Delete(id int) error {
	result := repository.DB.Delete(&domain.Teacher{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
