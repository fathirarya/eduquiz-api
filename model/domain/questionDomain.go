package domain

type Question struct {
	ID       uint
	QuizID   uint
	Quiz     Quiz
	Question string
	Option1  string
	Option2  string
	Option3  string
	Option4  string
	Title    string
}
