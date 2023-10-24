package res

import (
	"eduquiz-api/model/domain"
	"eduquiz-api/model/schema"
	"eduquiz-api/model/web"
)

func UserSchemaToUserDomain(user *schema.Users) *domain.Users {
	return &domain.Users{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Fullname: user.Fullname,
		Roles:    user.Roles,
	}
}

func UserDomainToUserResponse(user *domain.Users) web.UserResponse {
	return web.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Fullname: user.Fullname,
		Roles:    user.Roles,
	}
}

func UserDomainToUserLoginResponse(user *domain.Users) web.UserLoginResponse {
	return web.UserLoginResponse{
		Fullname: user.Fullname,
		Email:    user.Email,
		Roles:    user.Roles,
	}
}

func ConvertUserResponse(users []domain.Users) []web.UserResponse {
	var results []web.UserResponse
	for _, user := range users {
		userResponse := web.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Fullname: user.Fullname,
			Roles:    user.Roles,
		}
		results = append(results, userResponse)
	}
	return results
}
