package web

type KeyAnswerResponse struct {
	ID         uint   `json:"id"`
	QuestionID uint   `json:"question_id"`
	Questions  string `json:"questions"`
	Answer     string `json:"answer"`
}
