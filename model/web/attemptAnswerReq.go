package web

type AttemptAnswerCreateReq struct {
	QuestionID uint   `json:"question_id" validate:"required"`
	Answer     string `json:"answer" validate:"required,min=1"`
}
