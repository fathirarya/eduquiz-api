package web

import "eduquiz-api/model/domain"

type QuizResponse struct {
	ID             uint                `json:"id"`
	Title          string              `json:"title"`
	Description    string              `json:"description"`
	QuizCategoryID uint                `json:"quiz_category_id"`
	QuizCategory   domain.QuizCategory `json:"quiz_category"`
}
