package web

type QuizResultResponse struct {
	ID        uint   `json:"id"`
	StudentID uint   `json:"student_id"`
	FullName  string `json:"full_name"`
	QuizID    uint   `json:"quiz_id"`
	Title     string `json:"title"`
	Score     uint   `json:"score"`
}
