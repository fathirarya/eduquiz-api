package web

import "time"

type QuizResponse struct {
	ID             uint      `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	TeacherID      uint      `json:"teacher_id"`
	FullName       string    `json:"full_name"`
	QuizCategoryID uint      `json:"quiz_category_id"`
	Category       string    `json:"category"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
}

type ImplementationOpenAiResponse struct {
	OutputMessage string `json:"output_message"`
}
