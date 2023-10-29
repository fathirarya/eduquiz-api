package web

type QuizResponse struct {
	ID             uint   `json:"id"`
	QuizCategoryID uint   `json:"quiz_category_id"`
	Category       string `json:"category"`
	Title          string `json:"title"`
	Description    string `json:"description"`
}
