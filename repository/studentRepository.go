package repository

import (
	"eduquiz-api/model/domain"
	"eduquiz-api/model/schema"
	"eduquiz-api/utils/req"
	"eduquiz-api/utils/res"

	"gorm.io/gorm"
)

type StudentRepository interface {
	Create(user *domain.Student) (*domain.Student, error)
	Update(user *domain.Student, id int) (*domain.Student, error)
	FindById(id int) (*domain.Student, error)
	FindByEmail(email string) (*domain.Student, error)
	FindAll() ([]domain.Student, error)
	FindByUsername(name string) (*domain.Student, error)
	Delete(id int) error
}

type StudentRepositoryImpl struct {
	DB *gorm.DB
}

func NewStudentRepository(DB *gorm.DB) StudentRepository {
	return &StudentRepositoryImpl{DB: DB}
}

func (repository *StudentRepositoryImpl) Create(student *domain.Student) (*domain.Student, error) {
	studentDb := req.StudentDomainToStudentSchema(*student)
	result := repository.DB.Create(&studentDb)
	if result.Error != nil {
		return nil, result.Error
	}

	results := res.StudentSchemaToStudentDomain(studentDb)

	return results, nil
}

func (repository *StudentRepositoryImpl) Update(student *domain.Student, id int) (*domain.Student, error) {
	result := repository.DB.Table("students").Where("id = ?", id).Updates(domain.Student{Fullname: student.Fullname, Email: student.Email, Password: student.Password})
	if result.Error != nil {
		return nil, result.Error
	}

	return student, nil
}

func (repository *StudentRepositoryImpl) FindById(id int) (*domain.Student, error) {
	student := domain.Student{}

	result := repository.DB.First(&student, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &student, nil
}

func (repository *StudentRepositoryImpl) FindByEmail(email string) (*domain.Student, error) {
	student := domain.Student{}

	result := repository.DB.Where("email = ?", email).First(&student)
	if result.Error != nil {
		return nil, result.Error
	}

	return &student, nil
}

func (repository *StudentRepositoryImpl) FindAll() ([]domain.Student, error) {
	student := []domain.Student{}

	query := `SELECT * FROM students WHERE delete_at IS NULL`
	result := repository.DB.Raw(query).Scan(&student)
	if result.Error != nil {
		return nil, result.Error
	}

	return student, nil
}

func (repository *StudentRepositoryImpl) FindByUsername(username string) (*domain.Student, error) {
	student := domain.Student{}

	// Menggunakan query LIKE yang tidak case-sensitive
	result := repository.DB.Where("LOWER(username) LIKE LOWER(?)", "%"+username+"%").First(&student)

	if result.Error != nil {
		return nil, result.Error
	}

	return &student, nil
}

func (repository *StudentRepositoryImpl) Delete(id int) error {
	result := repository.DB.Delete(&schema.Student{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
