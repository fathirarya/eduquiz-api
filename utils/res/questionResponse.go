package res

import (
	"eduquiz-api/model/domain"
	"eduquiz-api/model/schema"
	"eduquiz-api/model/web"
)

func QuestionSchemaToQuestionDomain(question *schema.Question) *domain.Question {
	return &domain.Question{
		ID:       question.ID,
		QuizID:   question.QuizID,
		Question: question.Question,
		Option1:  question.Option1,
		Option2:  question.Option2,
		Option3:  question.Option3,
		Option4:  question.Option4,
	}
}

func QuestionDomainToQuestionResponse(question *domain.Question) *web.QuestionResponse {
	return &web.QuestionResponse{
		ID:       question.ID,
		QuizID:   question.QuizID,
		Quiz:     question.Quiz,
		Question: question.Question,
		Option1:  question.Option1,
		Option2:  question.Option2,
		Option3:  question.Option3,
		Option4:  question.Option4,
	}
}

func ConvertQuestionResponse(questions []domain.Question) []web.QuestionResponse {
	var results []web.QuestionResponse
	for _, question := range questions {
		questionResponse := web.QuestionResponse{
			ID:       question.ID,
			QuizID:   question.QuizID,
			Quiz:     question.Quiz,
			Question: question.Question,
			Option1:  question.Option1,
			Option2:  question.Option2,
			Option3:  question.Option3,
			Option4:  question.Option4,
		}
		results = append(results, questionResponse)
	}
	return results
}
