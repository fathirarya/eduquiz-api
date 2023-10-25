package web

type QuizResponse struct {
	ID             uint   `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	QuizCategoryID uint   `json:"quiz_category_id"`
}
