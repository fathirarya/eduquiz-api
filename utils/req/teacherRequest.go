package req

import (
	"eduquiz-api/model/domain"
	"eduquiz-api/model/schema"
	"eduquiz-api/model/web"
)

func TeacherCreateRequestToTeacherDomain(request web.TeacherCreateRequest) *domain.Teacher {
	return &domain.Teacher{
		Username: request.Username,
		Password: request.Password,
		Email:    request.Email,
		Fullname: request.Fullname,
	}
}

func TeacherUpdateRequestToTeacherDomain(request web.TeacherUpdateRequest) *domain.Teacher {
	return &domain.Teacher{
		Username: request.Username,
		Password: request.Password,
		Email:    request.Email,
		Fullname: request.Fullname,
	}
}

func TeacherLoginRequestToTeacherDomain(request web.TeacherLoginRequest) *domain.Teacher {
	return &domain.Teacher{
		Email:    request.Email,
		Password: request.Password,
	}
}

func TeacherDomainToTeacherSchema(teacher domain.Teacher) *schema.Teacher {
	return &schema.Teacher{
		ID:       teacher.ID,
		Username: teacher.Username,
		Password: teacher.Password,
		Email:    teacher.Email,
		Fullname: teacher.Fullname,
	}
}
