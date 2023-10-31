package req

import (
	"eduquiz-api/model/domain"
	"eduquiz-api/model/schema"
	"eduquiz-api/model/web"
)

func QuizResultRequestToQuizResultDomain(request web.QuizResultCreateRequest) *domain.QuizResult {
	return &domain.QuizResult{
		StudentID: request.StudentID,
		QuizID:    request.QuizID,
	}
}

func QuizResultDomainToQuizResultSchema(quizResult domain.QuizResult) *schema.QuizResult {
	return &schema.QuizResult{
		StudentID: quizResult.StudentID,
		QuizID:    quizResult.QuizID,
	}
}
