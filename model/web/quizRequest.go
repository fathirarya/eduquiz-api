package web

type QuizCreateRequest struct {
	Title          string `json:"title" validate:"required,min=1,max=255"`
	Description    string `json:"description" validate:"required,min=1,max=255"`
	QuizCategoryID uint   `json:"quiz_category_id" validate:"required"`
}

type QuizUpdateRequest struct {
	Title       string `json:"title" validate:"required,min=1,max=255"`
	Description string `json:"description" validate:"required,min=1,max=255"`
}
