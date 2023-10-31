package repository

import (
	"eduquiz-api/model/domain"
	"eduquiz-api/utils/req"
	"eduquiz-api/utils/res"

	"gorm.io/gorm"
)

type AttemptAnswerRepository interface {
	PostAnswer(attemptAnswer *domain.AttemptAnswer) (*domain.AttemptAnswer, error)
	FindById(questionId int) (*domain.AttemptAnswer, error)
	FindAll() ([]domain.AttemptAnswer, error)
	FindByStudentId(studentId int) ([]domain.AttemptAnswer, error)
}

type AttemptAnswerRepositoryImpl struct {
	DB *gorm.DB
}

func NewAttemptAnswerRepository(DB *gorm.DB) AttemptAnswerRepository {
	return &AttemptAnswerRepositoryImpl{DB: DB}
}

func (repository *AttemptAnswerRepositoryImpl) PostAnswer(attemptAnswer *domain.AttemptAnswer) (*domain.AttemptAnswer, error) {
	attemptAnswerDb := req.AttemptAnswerDomainToAttemptAnswerSchema(*attemptAnswer)
	result := repository.DB.Create(&attemptAnswer)
	if result.Error != nil {
		return nil, result.Error
	}

	results := res.AttemptAnswerSchemaToAttemptAnswerDomain(attemptAnswerDb)

	return results, nil

}

func (repository *AttemptAnswerRepositoryImpl) FindById(id int) (*domain.AttemptAnswer, error) {
	questionDb := domain.AttemptAnswer{}

	if err := repository.DB.First(&questionDb, id).Error; err != nil {
		return nil, err
	}

	query := `SELECT attempt_answers.*, 
	FROM attempt_answers
	LEFT JOIN questions ON attempt_answers.question_id = questions.id
	LEFT JOIN students ON attempt_answers.student_id = students.id
	LEFT JOIN quizzes ON attempt_answers.quiz_id = quizzes.id
	WHERE attempt_answers.id = ?`

	result := repository.DB.Raw(query, id).Scan(&questionDb)

	if result.Error != nil {
		return nil, result.Error

	}

	return &questionDb, nil
}

func (repository *AttemptAnswerRepositoryImpl) FindAll() ([]domain.AttemptAnswer, error) {
	var attemptAnswers []domain.AttemptAnswer

	query := `SELECT attempt_answers.*, questions.question AS quest
	FROM attempt_answers
	LEFT JOIN questions ON attempt_answers.question_id = questions.id
	LEFT JOIN students ON attempt_answers.student_id = students.id
	LEFT JOIN quizzes ON attempt_answers.quiz_id = quizzes.id`

	result := repository.DB.Raw(query).Scan(&attemptAnswers)

	if result.Error != nil {
		return nil, result.Error

	}

	return attemptAnswers, nil
}

func (repository *AttemptAnswerRepositoryImpl) FindByStudentId(studentId int) ([]domain.AttemptAnswer, error) {
	var attemptAnswers []domain.AttemptAnswer

	query := `SELECT attempt_answers.*, questions.question AS quest
	FROM attempt_answers
	LEFT JOIN questions ON attempt_answers.question_id = questions.id
	LEFT JOIN students ON attempt_answers.student_id = students.id
	LEFT JOIN quizzes ON attempt_answers.quiz_id = quizzes.id
	WHERE attempt_answers.student_id = ?`

	result := repository.DB.Raw(query, studentId).Scan(&attemptAnswers)

	if result.Error != nil {
		return nil, result.Error

	}

	return attemptAnswers, nil
}
