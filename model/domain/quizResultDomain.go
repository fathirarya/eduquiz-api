package domain

type QuizResult struct {
	ID        uint
	StudentID uint
	Student   Student
	QuizID    uint
	Quiz      Quiz
	Score     int
}
