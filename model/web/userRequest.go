package web

type UserCreateRequest struct {
	Username string `json:"username" validate:"required,min=1,max=255"`
	Password string `json:"password" validate:"required,min=8,max=255"`
	Email    string `json:"email" validate:"required,email,min=1,max=255"`
	Fullname string `json:"fullname" validate:"required,min=1,max=255"`
	Roles    string `json:"roles" validate:"required,min=1,max=255"`
}

type UserUpdateRequest struct {
	Username string `json:"username" validate:"required,min=1,max=255"`
	Password string `json:"password" validate:"required,min=8,max=255"`
	Email    string `json:"email" validate:"required,email,min=1,max=255"`
	Fullname string `json:"fullname" validate:"required,min=1,max=255"`
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email,min=1,max=255"`
	Password string `json:"password" validate:"required,min=8,max=255"`
}
