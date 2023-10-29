package web

type TeacherResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
}

type TeacherLoginResponse struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
