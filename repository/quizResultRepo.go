package repository

import (
	"eduquiz-api/model/domain"
	"eduquiz-api/utils/req"
	"eduquiz-api/utils/res"

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
	quizResultDb := req.QuizResultDomainToQuizResultSchema(*quizResult)
	result := repository.DB.Create(&quizResultDb)
	if result.Error != nil {
		return nil, result.Error
	}

	results := res.QuizResultSchemaToQuizResultDomain(quizResultDb)

	return results, nil
}
