package repository

import (
	"eduquiz-api/model/domain"
	"eduquiz-api/utils/req"
	"eduquiz-api/utils/res"

	"gorm.io/gorm"
)

type KeyAnswerRepository interface {
	Create(keyAnswer *domain.KeyAnswer) (*domain.KeyAnswer, error)
	Update(keyAnswer *domain.KeyAnswer, id int) (*domain.KeyAnswer, error)
	FindById(id int) (*domain.KeyAnswer, error)
	FindAll() ([]domain.KeyAnswer, error)
	Delete(id int) error
}

type KeyAnswerRepositoryImpl struct {
	DB *gorm.DB
}

func NewKeyAnswerRepository(DB *gorm.DB) KeyAnswerRepository {
	return &KeyAnswerRepositoryImpl{DB: DB}
}

func (repository *KeyAnswerRepositoryImpl) Create(keyAnswer *domain.KeyAnswer) (*domain.KeyAnswer, error) {
	keyAnswerDb := req.KeyAnswerDomainToKeyAnswerSchema(*keyAnswer)
	result := repository.DB.Create(&keyAnswerDb)
	if result.Error != nil {
		return nil, result.Error
	}

	results := res.KeyAnswerSchemaToKeyAnswerDomain(keyAnswerDb)

	return results, nil
}

func (repository *KeyAnswerRepositoryImpl) Update(keyAnswer *domain.KeyAnswer, id int) (*domain.KeyAnswer, error) {
	result := repository.DB.Table("key_answers").Where("id = ?", id).Updates(domain.KeyAnswer{QuestionID: keyAnswer.QuestionID, Answer: keyAnswer.Answer})
	if result.Error != nil {
		return nil, result.Error
	}

	return keyAnswer, nil
}

func (repository *KeyAnswerRepositoryImpl) FindById(id int) (*domain.KeyAnswer, error) {
	keyAnswer := domain.KeyAnswer{}

	if err := repository.DB.First(&keyAnswer, id).Error; err != nil {
		return nil, err
	}

	query := `SELECT key_answers.*, questions.question AS questions
	FROM key_answers
	LEFT JOIN questions ON key_answers.question_id = questions.id
	WHERE key_answers.id = ?`

	result := repository.DB.Raw(query, id).Scan(&keyAnswer)

	if result.Error != nil {
		return nil, result.Error
	}

	return &keyAnswer, nil
}

func (repository *KeyAnswerRepositoryImpl) FindAll() ([]domain.KeyAnswer, error) {
	keyAnswers := []domain.KeyAnswer{}
	query := `SELECT key_answers.*, questions.question AS questions
	FROM key_answers
	INNER JOIN questions ON key_answers.question_id = questions.id`

	result := repository.DB.Raw(query).Scan(&keyAnswers)

	if result.Error != nil {
		return nil, result.Error
	}

	return keyAnswers, nil
}

func (repository *KeyAnswerRepositoryImpl) Delete(id int) error {
	result := repository.DB.Delete(&domain.KeyAnswer{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
