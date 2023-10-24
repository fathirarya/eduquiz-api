package web

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Roles    string `json:"roles"`
}

type UserLoginResponse struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Roles    string `json:"roles"`
	Token    string `json:"token"`
}
