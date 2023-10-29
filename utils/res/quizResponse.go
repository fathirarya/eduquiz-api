package res

import (
	"eduquiz-api/model/domain"
	"eduquiz-api/model/schema"
	"eduquiz-api/model/web"
)

func QuizSchemaToQuizDomain(quiz *schema.Quiz) *domain.Quiz {
	return &domain.Quiz{
		ID:             quiz.ID,
		QuizCategoryID: quiz.QuizCategoryID,
		Title:          quiz.Title,
		Description:    quiz.Description,
	}
}

func QuizDomainToQuizResponse(quiz *domain.Quiz) web.QuizResponse {
	return web.QuizResponse{
		ID:             quiz.ID,
		QuizCategoryID: quiz.QuizCategoryID,
		Category:       quiz.Category,
		Title:          quiz.Title,
		Description:    quiz.Description,
	}
}

func ConvertQuizResponse(quizzes []domain.Quiz) []web.QuizResponse {
	var results []web.QuizResponse
	for _, quiz := range quizzes {
		quizResponse := web.QuizResponse{
			ID:             quiz.ID,
			QuizCategoryID: quiz.QuizCategoryID,
			Category:       quiz.Category,
			Title:          quiz.Title,
			Description:    quiz.Description,
		}
		results = append(results, quizResponse)
	}
	return results
}

func FindQuizToQuizResponse(quiz *domain.Quiz) web.QuizResponse {
	return web.QuizResponse{
		ID:             quiz.ID,
		QuizCategoryID: quiz.QuizCategoryID,
		Category:       quiz.Category,
		Title:          quiz.Title,
		Description:    quiz.Description,
	}
}
