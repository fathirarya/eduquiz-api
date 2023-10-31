package web

type AttemptAnswerResponse struct {
	ID         uint   `json:"id"`
	StudentID  uint   `json:"student_id"`
	QuizID     uint   `json:"quiz_id"`
	QuestionID uint   `json:"question_id"`
	Answer     string `json:"answer"`
	IsCorrect  bool   `json:"is_correct"`
}
