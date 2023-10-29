package web

type TeacherCreateRequest struct {
	Username string `json:"username" validate:"required,min=1,max=255"`
	Password string `json:"password" validate:"required,min=8,max=255"`
	Email    string `json:"email" validate:"required,email,min=1,max=255"`
	Fullname string `json:"fullname" validate:"required,min=1,max=255"`
}

type TeacherUpdateRequest struct {
	Username string `json:"username" validate:"required,min=1,max=255"`
	Password string `json:"password" validate:"required,min=8,max=255"`
	Email    string `json:"email" validate:"required,email,min=1,max=255"`
	Fullname string `json:"fullname" validate:"required,min=1,max=255"`
}

type TeacherLoginRequest struct {
	Email    string `json:"email" validate:"required,email,min=1,max=255"`
	Password string `json:"password" validate:"required,min=8,max=255"`
}
