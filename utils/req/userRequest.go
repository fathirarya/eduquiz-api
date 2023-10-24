package req

import (
	"eduquiz-api/model/domain"
	"eduquiz-api/model/schema"
	"eduquiz-api/model/web"
)

func UserCreateRequestToUserDomain(request web.UserCreateRequest) *domain.Users {
	return &domain.Users{
		Username: request.Username,
		Password: request.Password,
		Email:    request.Email,
		Fullname: request.Fullname,
		Roles:    request.Roles,
	}
}

func UserUpdateRequestToUserDomain(request web.UserUpdateRequest) *domain.Users {
	return &domain.Users{
		Username: request.Username,
		Password: request.Password,
		Email:    request.Email,
		Fullname: request.Fullname,
	}
}

func UserLoginRequestToUserDomain(request web.UserLoginRequest) *domain.Users {
	return &domain.Users{
		Email:    request.Email,
		Password: request.Password,
	}
}

func UserDomainToUserSchema(user domain.Users) *schema.Users {
	return &schema.Users{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Fullname: user.Fullname,
		Roles:    user.Roles,
	}
}
