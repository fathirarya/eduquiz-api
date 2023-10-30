package req

import (
	"eduquiz-api/model/domain"
	"eduquiz-api/model/schema"
	"eduquiz-api/model/web"
)

func AttemptAnswerCreateRequestToAttemptAnswerDomain(request web.AttemptAnswerCreateReq) *domain.AttemptAnswer {
	return &domain.AttemptAnswer{
		QuestionID: request.QuestionID,
		Answer:     request.Answer,
	}
}

func AttemptAnswerDomainToAttemptAnswerSchema(attemptAnswer domain.AttemptAnswer) *schema.AttemptAnswer {
	return &schema.AttemptAnswer{
		QuestionID: attemptAnswer.QuestionID,
		Answer:     attemptAnswer.Answer,
	}
}
