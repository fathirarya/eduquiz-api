package web

type QuizResultResponse struct {
	ID        uint `json:"id"`
	StudentID uint `json:"student_id"`
	QuizID    uint `json:"quiz_id"`
	Score     int  `json:"score"`
}
