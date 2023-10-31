package res

import (
	"eduquiz-api/model/domain"
	"eduquiz-api/model/schema"
	"eduquiz-api/model/web"
)

func AttemptAnswerSchemaToAttemptAnswerDomain(attemptAnswer *schema.AttemptAnswer) *domain.AttemptAnswer {
	return &domain.AttemptAnswer{
		ID:         attemptAnswer.ID,
		StudentID:  attemptAnswer.StudentID,
		QuizID:     attemptAnswer.QuizID,
		QuestionID: attemptAnswer.QuestionID,
		Answer:     attemptAnswer.Answer,
		IsCorrect:  attemptAnswer.IsCorrect,
	}
}

func AttemptAnswerDomainToAttemptAnswerResponse(attemptAnswer *domain.AttemptAnswer) *web.AttemptAnswerResponse {
	return &web.AttemptAnswerResponse{
		ID:         attemptAnswer.ID,
		StudentID:  attemptAnswer.StudentID,
		QuizID:     attemptAnswer.QuizID,
		QuestionID: attemptAnswer.QuestionID,
		Answer:     attemptAnswer.Answer,
		IsCorrect:  attemptAnswer.IsCorrect,
	}
}

func ConvertAttemptAnswerResponse(attemptAnswers []domain.AttemptAnswer) []web.AttemptAnswerResponse {
	var results []web.AttemptAnswerResponse
	for _, attemptAnswer := range attemptAnswers {
		attemptAnswerResponse := web.AttemptAnswerResponse{
			ID:         attemptAnswer.ID,
			StudentID:  attemptAnswer.StudentID,
			QuizID:     attemptAnswer.QuizID,
			QuestionID: attemptAnswer.QuestionID,
			Answer:     attemptAnswer.Answer,
			IsCorrect:  attemptAnswer.IsCorrect,
		}
		results = append(results, attemptAnswerResponse)
	}
	return results
}

func FindAttemptAnswerDomainToAttemptAnswerResponse(attemptAnswer *domain.AttemptAnswer) *web.AttemptAnswerResponse {
	return &web.AttemptAnswerResponse{
		ID:         attemptAnswer.ID,
		StudentID:  attemptAnswer.StudentID,
		QuizID:     attemptAnswer.QuizID,
		QuestionID: attemptAnswer.QuestionID,
		Answer:     attemptAnswer.Answer,
		IsCorrect:  attemptAnswer.IsCorrect,
	}
}
