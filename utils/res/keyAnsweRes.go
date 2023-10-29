package res

import (
	"eduquiz-api/model/domain"
	"eduquiz-api/model/schema"
	"eduquiz-api/model/web"
)

func KeyAnswerSchemaToKeyAnswerDomain(keyAnswer *schema.KeyAnswer) *domain.KeyAnswer {
	return &domain.KeyAnswer{
		ID:         keyAnswer.ID,
		QuestionID: keyAnswer.QuestionID,
		Answer:     keyAnswer.Answer,
	}
}

func KeyAnswerDomainToKeyAnswerResponse(keyAnswer *domain.KeyAnswer) *web.KeyAnswerResponse {
	return &web.KeyAnswerResponse{
		ID:         keyAnswer.ID,
		QuestionID: keyAnswer.QuestionID,
		Answer:     keyAnswer.Answer,
	}
}

func ConvertKeyAnswerResponse(keyAnswers []domain.KeyAnswer) []web.KeyAnswerResponse {
	var results []web.KeyAnswerResponse
	for _, keyAnswer := range keyAnswers {
		keyAnswerResponse := web.KeyAnswerResponse{
			ID:         keyAnswer.ID,
			QuestionID: keyAnswer.QuestionID,
			Questions:  keyAnswer.Questions,
			Answer:     keyAnswer.Answer,
		}
		results = append(results, keyAnswerResponse)
	}
	return results
}

func FindKeyAnswerToKeyAnswerResponse(keyAnswer *domain.KeyAnswer) *web.KeyAnswerResponse {
	return &web.KeyAnswerResponse{
		ID:         keyAnswer.ID,
		QuestionID: keyAnswer.QuestionID,
		Questions:  keyAnswer.Questions,
		Answer:     keyAnswer.Answer,
	}
}
