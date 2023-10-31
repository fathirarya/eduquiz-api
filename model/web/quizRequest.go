package web

type QuizCreateRequest struct {
	Title          string `json:"title" validate:"required,min=1,max=255"`
	Description    string `json:"description" validate:"required,min=1,max=255"`
	QuizCategoryID uint   `json:"quiz_category_id" validate:"required"`
	TeacherID      uint   `json:"teacher_id" validate:"required"`
}

type QuizUpdateRequest struct {
	Title       string `json:"title" validate:"required,min=1,max=255"`
	Description string `json:"description" validate:"required,min=1,max=255"`
}

type ImplementationOpenAiRequest struct {
	InputMessage string `json:"input_message" validate:"required,min=1,max=255"`
}
