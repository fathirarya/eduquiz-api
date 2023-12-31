package web

type QuizCategoryCreateRequest struct {
	ID       uint   `json:"id"`
	Category string `json:"category" validate:"required,min=1,max=255"`
}
