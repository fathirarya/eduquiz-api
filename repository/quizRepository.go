package repository

import (
	"eduquiz-api/model/domain"
	"eduquiz-api/utils/req"
	"eduquiz-api/utils/res"

	"gorm.io/gorm"
)

type QuizRepository interface {
	Create(quiz *domain.Quiz) (*domain.Quiz, error)
	Update(quiz *domain.Quiz, id int) (*domain.Quiz, error)
	FindByTitle(title string) (*domain.Quiz, error)
	FindById(id int) (*domain.Quiz, error)
	FindAll() ([]domain.Quiz, error)
	Delete(id int) error
}

type QuizRepositoryImpl struct {
	DB *gorm.DB
}

func NewQuizRepository(DB *gorm.DB) QuizRepository {
	return &QuizRepositoryImpl{DB: DB}
}

func (repository *QuizRepositoryImpl) Create(quiz *domain.Quiz) (*domain.Quiz, error) {
	quizDb := req.QuizDomainToQuizSchema(*quiz)
	result := repository.DB.Create(&quizDb)
	if result.Error != nil {
		return nil, result.Error
	}

	results := res.QuizSchemaToQuizDomain(quizDb)

	return results, nil
}

func (repository *QuizRepositoryImpl) Update(quiz *domain.Quiz, id int) (*domain.Quiz, error) {
	result := repository.DB.Table("quiz").Where("id = ?", id).Updates(domain.Quiz{Title: quiz.Title, Description: quiz.Description})
	if result.Error != nil {
		return nil, result.Error
	}

	return quiz, nil
}

func (repository *QuizRepositoryImpl) FindByTitle(title string) (*domain.Quiz, error) {
	quiz := domain.Quiz{}

	result := repository.DB.Where("title = ?", title).First(&quiz)
	if result.Error != nil {
		return nil, result.Error
	}

	return &quiz, nil
}

func (repository *QuizRepositoryImpl) FindById(id int) (*domain.Quiz, error) {
	quiz := domain.Quiz{}

	result := repository.DB.First(&quiz, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &quiz, nil
}

func (repository *QuizRepositoryImpl) FindAll() ([]domain.Quiz, error) {
	var quizzes []domain.Quiz

	result := repository.DB.Find(&quizzes)
	if result.Error != nil {
		return nil, result.Error
	}

	return quizzes, nil
}

func (repository *QuizRepositoryImpl) Delete(id int) error {
	result := repository.DB.Delete(&domain.Quiz{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
