package req

import (
	"eduquiz-api/model/domain"
	"eduquiz-api/model/schema"
	"eduquiz-api/model/web"
)

func QuizCategoryCreateRequestToQuizCategoryDomain(request web.QuizCategoryCreateRequest) *domain.QuizCategory {
	return &domain.QuizCategory{
		ID:       request.ID,
		Category: request.Category,
	}
}

func QuizCategoryDomainToQuizCategorySchema(QuizCategory domain.QuizCategory) *schema.QuizCategory {
	return &schema.QuizCategory{
		ID:       QuizCategory.ID,
		Category: QuizCategory.Category,
	}
}
