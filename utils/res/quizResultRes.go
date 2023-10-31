package res

import (
	"eduquiz-api/model/domain"
	"eduquiz-api/model/schema"
	"eduquiz-api/model/web"
)

func QuizResultSchemaToQuizResultDomain(quizResult *schema.QuizResult) *domain.QuizResult {
	return &domain.QuizResult{
		StudentID: quizResult.StudentID,
		QuizID:    quizResult.QuizID,
		Score:     quizResult.Score,
	}
}

func QuizResultDomainToQuizResultResponse(quizResult *domain.QuizResult) web.QuizResultResponse {
	return web.QuizResultResponse{
		StudentID: quizResult.StudentID,
		QuizID:    quizResult.QuizID,
		Score:     quizResult.Score,
	}
}

func ConvertQuizResultResponse(quizResults []domain.QuizResult) []web.QuizResultResponse {
	var results []web.QuizResultResponse
	for _, quizResult := range quizResults {
		quizResultResponse := web.QuizResultResponse{
			StudentID: quizResult.StudentID,
			QuizID:    quizResult.QuizID,
			Score:     quizResult.Score,
		}
		results = append(results, quizResultResponse)
	}
	return results
}
