package repository

import (
	"eduquiz-api/model/domain"
	"eduquiz-api/utils/req"
	"eduquiz-api/utils/res"

	"gorm.io/gorm"
)

type QuestionRepository interface {
	Create(question *domain.Question) (*domain.Question, error)
	Update(question *domain.Question, id int) (*domain.Question, error)
	FindByQuestion(question string) (*domain.Question, error)
	FindById(id int) (*domain.Question, error)
	FindAll() ([]domain.Question, error)
	Delete(id int) error
}

type QuestionRepositoryImpl struct {
	DB *gorm.DB
}

func NewQuestionRepository(DB *gorm.DB) QuestionRepository {
	return &QuestionRepositoryImpl{DB: DB}
}

func (repository *QuestionRepositoryImpl) Create(question *domain.Question) (*domain.Question, error) {
	questionDb := req.QuestionDomainToQuestionSchema(*question)
	result := repository.DB.Create(&questionDb)
	if result.Error != nil {
		return nil, result.Error
	}

	results := res.QuestionSchemaToQuestionDomain(questionDb)

	return results, nil
}

func (repository *QuestionRepositoryImpl) Update(question *domain.Question, id int) (*domain.Question, error) {
	result := repository.DB.Table("questions").Where("id = ?", id).Updates(domain.Question{Question: question.Question, Option1: question.Option1, Option2: question.Option2, Option3: question.Option3, Option4: question.Option4})
	if result.Error != nil {
		return nil, result.Error
	}

	return question, nil
}

func (repository *QuestionRepositoryImpl) FindByQuestion(question string) (*domain.Question, error) {
	questionDb := domain.Question{}

	result := repository.DB.Where("question = ?", question).First(&questionDb)
	if result.Error != nil {
		return nil, result.Error
	}

	return &questionDb, nil
}

func (repository *QuestionRepositoryImpl) FindById(id int) (*domain.Question, error) {
	var question domain.Question

	result := repository.DB.Preload("Quiz").First(&question, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &question, nil
}

func (repository *QuestionRepositoryImpl) FindAll() ([]domain.Question, error) {
	var questions []domain.Question

	if err := repository.DB.Find(&questions).Error; err != nil {
		return nil, err
	}

	result := repository.DB.Preload("Quiz").Find(&questions)

	if result.Error != nil {
		return nil, result.Error
	}

	return questions, nil
}

func (repository *QuestionRepositoryImpl) Delete(id int) error {
	result := repository.DB.Delete(&domain.Question{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
