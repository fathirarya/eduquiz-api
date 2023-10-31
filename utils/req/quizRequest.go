package req

import (
	"eduquiz-api/model/domain"
	"eduquiz-api/model/schema"
	"eduquiz-api/model/web"
)

func QuizCreateRequestToQuizDomain(request web.QuizCreateRequest) *domain.Quiz {
	return &domain.Quiz{
		Title:          request.Title,
		Description:    request.Description,
		QuizCategoryID: request.QuizCategoryID,
		TeacherID:      request.TeacherID,
	}
}

func QuizUpdateRequestToQuizDomain(request web.QuizUpdateRequest) *domain.Quiz {
	return &domain.Quiz{
		Title:       request.Title,
		Description: request.Description,
	}
}

func QuizDomainToQuizSchema(quiz domain.Quiz) *schema.Quiz {
	return &schema.Quiz{
		Title:          quiz.Title,
		Description:    quiz.Description,
		QuizCategoryID: quiz.QuizCategoryID,
		TeacherID:      quiz.TeacherID,
	}
}

func ImplementationOpenAiRequestToImplementationOpenAiDomain(request web.ImplementationOpenAiRequest) *domain.OpenAi {
	return &domain.OpenAi{
		InputMessage: request.InputMessage,
	}
}
