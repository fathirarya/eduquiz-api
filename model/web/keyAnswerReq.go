package web

type KeyCreateAnswerReq struct {
	QuestionID uint   `json:"question_id" validate:"required"`
	Answer     string `json:"answer" validate:"required,min=1"`
}

type KeyUpdateAnswerReq struct {
	QuestionID uint   `json:"question_id" validate:"required"`
	Answer     string `json:"answer" validate:"required,min=1"`
}
