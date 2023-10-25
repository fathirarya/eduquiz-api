package repository

import (
	"eduquiz-api/model/domain"
	"eduquiz-api/model/schema"
	"eduquiz-api/utils/req"
	"eduquiz-api/utils/res"
	"fmt"

	"gorm.io/gorm"
)

type QuizCategoryRepository interface {
	Create(quizCategory *domain.QuizCategory) (*domain.QuizCategory, error)
	FindById(id int) (*domain.QuizCategory, error)
	FindByName(name string) (*domain.QuizCategory, error)
	FindAll() ([]domain.QuizCategory, error)
	Delete(id int) error
}

type QuizCategoryRepositoryImpl struct {
	DB *gorm.DB
}

func NewQuizCategoryRepository(DB *gorm.DB) QuizCategoryRepository {
	return &QuizCategoryRepositoryImpl{DB: DB}
}

func (repository *QuizCategoryRepositoryImpl) Create(quizCategory *domain.QuizCategory) (*domain.QuizCategory, error) {
	quizCategoryDb := req.QuizCategoryDomainToQuizCategorySchema(*quizCategory)
	result := repository.DB.Create(&quizCategory)
	if result.Error != nil {
		return nil, result.Error
	}

	results := res.QuizCategorySchemaToQuizCategoryDomain(quizCategoryDb)
	fmt.Println(results)
	return results, nil
}

func (repository *QuizCategoryRepositoryImpl) FindById(id int) (*domain.QuizCategory, error) {
	quizCategory := domain.QuizCategory{}

	result := repository.DB.First(&quizCategory, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &quizCategory, nil
}

func (repository *QuizCategoryRepositoryImpl) FindByName(name string) (*domain.QuizCategory, error) {
	quizCategory := domain.QuizCategory{}

	result := repository.DB.Where("name = ?", name).First(&quizCategory)
	if result.Error != nil {
		return nil, result.Error
	}

	return &quizCategory, nil
}

func (repository *QuizCategoryRepositoryImpl) FindAll() ([]domain.QuizCategory, error) {
	quizCategory := []domain.QuizCategory{}

	result := repository.DB.Find(&quizCategory)
	if result.Error != nil {
		return nil, result.Error
	}

	return quizCategory, nil
}

func (repository *QuizCategoryRepositoryImpl) Delete(id int) error {
	result := repository.DB.Delete(&schema.QuizCategory{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
