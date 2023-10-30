package web

type AttemptAnswerResponse struct {
	ID         uint `json:"id"`
	QuestionID uint `json:"question_id"`
	// Quest      string `json:"quest"`
	Answer    string `json:"answer"`
	IsCorrect bool   `json:"is_correct"`
}
