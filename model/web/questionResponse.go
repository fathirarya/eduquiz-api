package web

type QuestionResponse struct {
	ID       uint   `json:"id"`
	QuizID   uint   `json:"quiz_id"`
	Title    string `json:"title"`
	Question string `json:"question"`
	Option1  string `json:"option_1"`
	Option2  string `json:"option_2"`
	Option3  string `json:"option_3"`
	Option4  string `json:"option_4"`
}
