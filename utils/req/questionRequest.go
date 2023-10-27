package req

import (
	"eduquiz-api/model/domain"
	"eduquiz-api/model/schema"
	"eduquiz-api/model/web"
)

func QuestionCreateRequestToQuestionDomain(request web.QuestionCreateRequest) *domain.Question {
	return &domain.Question{
		QuizID:   request.QuizID,
		Question: request.Question,
		Option1:  request.Option1,
		Option2:  request.Option2,
		Option3:  request.Option3,
		Option4:  request.Option4,
	}
}

func QuestionUpdateRequestToQuestionDomain(request web.QuestionUpdateRequest) *domain.Question {
	return &domain.Question{
		QuizID:   request.QuizID,
		Question: request.Question,
		Option1:  request.Option1,
		Option2:  request.Option2,
		Option3:  request.Option3,
	}
}

func QuestionDomainToQuestionSchema(question domain.Question) *schema.Question {
	return &schema.Question{
		QuizID:   question.QuizID,
		Question: question.Question,
		Option1:  question.Option1,
		Option2:  question.Option2,
		Option3:  question.Option3,
		Option4:  question.Option4,
	}
}
