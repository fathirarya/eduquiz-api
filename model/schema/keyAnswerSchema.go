package schema

type KeyAnswer struct {
	ID         uint     `json:"id" gorm:"primaryKey"`
	QuestionID uint     `json:"question_id"`
	Question   Question `json:"question" gorm:"foreignKey:QuestionID"`
	Answer     string   `json:"answer"`
}
