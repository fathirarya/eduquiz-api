package web

type AttemptAnswerCreateReq struct {
	StudentID  uint   `json:"student_id" validate:"required"`
	QuizID     uint   `json:"quiz_id" validate:"required"`
	QuestionID uint   `json:"question_id" validate:"required"`
	Answer     string `json:"answer" validate:"required,min=1"`
}
