package web

type QuizResultCreateRequest struct {
	StudentID uint `json:"student_id" validate:"required"`
	QuizID    uint `json:"quiz_id" validate:"required"`
}
