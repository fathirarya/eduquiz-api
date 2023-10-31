package req

import (
	"eduquiz-api/model/domain"
	"eduquiz-api/model/schema"
	"eduquiz-api/model/web"
)

func AttemptAnswerCreateRequestToAttemptAnswerDomain(request web.AttemptAnswerCreateReq) *domain.AttemptAnswer {
	return &domain.AttemptAnswer{
		StudentID:  request.StudentID,
		QuizID:     request.QuizID,
		QuestionID: request.QuestionID,
		Answer:     request.Answer,
	}
}

func AttemptAnswerDomainToAttemptAnswerSchema(attemptAnswer domain.AttemptAnswer) *schema.AttemptAnswer {
	return &schema.AttemptAnswer{
		StudentID:  attemptAnswer.StudentID,
		QuizID:     attemptAnswer.QuizID,
		QuestionID: attemptAnswer.QuestionID,
		Answer:     attemptAnswer.Answer,
	}
}
