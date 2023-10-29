package req

import (
	"eduquiz-api/model/domain"
	"eduquiz-api/model/schema"
	"eduquiz-api/model/web"
)

func StudentCreateRequestToStudentDomain(request web.StudentCreateRequest) *domain.Student {
	return &domain.Student{
		Username: request.Username,
		Password: request.Password,
		Email:    request.Email,
		Fullname: request.Fullname,
	}
}

func StudentUpdateRequestToStudentDomain(request web.StudentUpdateRequest) *domain.Student {
	return &domain.Student{
		Username: request.Username,
		Password: request.Password,
		Email:    request.Email,
		Fullname: request.Fullname,
	}
}

func StudentLoginRequestToStudentDomain(request web.StudentLoginRequest) *domain.Student {
	return &domain.Student{
		Email:    request.Email,
		Password: request.Password,
	}
}

func StudentDomainToStudentSchema(user domain.Student) *schema.Student {
	return &schema.Student{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Fullname: user.Fullname,
	}
}
