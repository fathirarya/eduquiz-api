package res

import (
	"eduquiz-api/model/domain"
	"eduquiz-api/model/schema"
	"eduquiz-api/model/web"
)

func QuizCategorySchemaToQuizCategoryDomain(quizCategory *schema.QuizCategory) *domain.QuizCategory {
	return &domain.QuizCategory{
		ID:       quizCategory.ID,
		Category: quizCategory.Category,
	}
}

func QuizCategoryDomainToQuizCategoryResponse(quizCategory *domain.QuizCategory) web.QuizCategoryResponse {
	return web.QuizCategoryResponse{
		Category: quizCategory.Category,
	}
}

func ConvertQuizCategoryResponse(quizCategory []domain.QuizCategory) []web.QuizCategoryResponse {
	var results []web.QuizCategoryResponse
	for _, quizCategory := range quizCategory {
		quizCategoryResponse := web.QuizCategoryResponse{
			Category: quizCategory.Category,
		}
		results = append(results, quizCategoryResponse)
	}
	return results
}
