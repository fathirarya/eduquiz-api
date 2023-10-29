package req

import (
	"eduquiz-api/model/domain"
	"eduquiz-api/model/schema"
	"eduquiz-api/model/web"
)

func KeyCreateAnswerRequestToKeyAnswerDomain(request web.KeyCreateAnswerReq) *domain.KeyAnswer {
	return &domain.KeyAnswer{
		QuestionID: request.QuestionID,
		Answer:     request.Answer,
	}
}

func KeyUpdateAnswerRequestToKeyAnswerDomain(request web.KeyUpdateAnswerReq) *domain.KeyAnswer {
	return &domain.KeyAnswer{
		QuestionID: request.QuestionID,
		Answer:     request.Answer,
	}
}

func KeyAnswerDomainToKeyAnswerSchema(keyAnswer domain.KeyAnswer) *schema.KeyAnswer {
	return &schema.KeyAnswer{
		QuestionID: keyAnswer.QuestionID,
		Answer:     keyAnswer.Answer,
	}
}
