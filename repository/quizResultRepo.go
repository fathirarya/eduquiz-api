package repository

import (
	"eduquiz-api/model/domain"

	"gorm.io/gorm"
)

type QuizResultRepository interface {
	PostResult(quizResult *domain.QuizResult) (*domain.QuizResult, error)
}

type QuizResultRepositoryImpl struct {
	DB *gorm.DB
}

func NewQuizResultRepository(DB *gorm.DB) QuizResultRepository {
	return &QuizResultRepositoryImpl{DB: DB}
}

func (repository *QuizResultRepositoryImpl) PostResult(quizResult *domain.QuizResult) (*domain.QuizResult, error) {
	result := repository.DB.Create(&quizResult)
	if result.Error != nil {
		return nil, result.Error
	}

	return quizResult, nil
}
