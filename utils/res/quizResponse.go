package res

import (
	"eduquiz-api/model/domain"
	"eduquiz-api/model/schema"
	"eduquiz-api/model/web"
)

func QuizSchemaToQuizDomain(quiz *schema.Quiz) *domain.Quiz {
	return &domain.Quiz{
		ID:             quiz.ID,
		Title:          quiz.Title,
		Description:    quiz.Description,
		QuizCategoryID: quiz.QuizCategoryID,
	}
}

func QuizDomainToQuizResponse(quiz *domain.Quiz) *web.QuizResponse {
	return &web.QuizResponse{
		ID:             quiz.ID,
		Title:          quiz.Title,
		Description:    quiz.Description,
		QuizCategoryID: quiz.QuizCategoryID,
	}
}

func ConvertQuizResponse(quizzes []domain.Quiz) []web.QuizResponse {
	var results []web.QuizResponse
	for _, quiz := range quizzes {
		quizResponse := web.QuizResponse{
			ID:             quiz.ID,
			Title:          quiz.Title,
			Description:    quiz.Description,
			QuizCategoryID: quiz.QuizCategoryID,
		}
		results = append(results, quizResponse)
	}
	return results
}
