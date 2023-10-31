package res

import (
	"eduquiz-api/model/domain"
	"eduquiz-api/model/schema"
	"eduquiz-api/model/web"
)

func TeacherSchemaToTeacherDomain(teacher *schema.Teacher) *domain.Teacher {
	return &domain.Teacher{
		ID:       teacher.ID,
		Username: teacher.Username,
		Password: teacher.Password,
		Email:    teacher.Email,
		Fullname: teacher.Fullname,
	}
}

func TeacherDomainToTeacherResponse(teacher *domain.Teacher) web.TeacherResponse {
	return web.TeacherResponse{
		ID:        teacher.ID,
		CreatedAt: teacher.CreatedAt,
		UpdatedAt: teacher.UpdatedAt,
		Username:  teacher.Username,
		Email:     teacher.Email,
		Fullname:  teacher.Fullname,
	}
}

func TeacherDomainToTeacherLoginResponse(teacher *domain.Teacher) web.TeacherLoginResponse {
	return web.TeacherLoginResponse{
		Fullname: teacher.Fullname,
		Email:    teacher.Email,
	}
}

func ConvertTeacherResponse(teachers []domain.Teacher) []web.TeacherResponse {
	var results []web.TeacherResponse
	for _, teacher := range teachers {
		teacherResponse := web.TeacherResponse{
			ID:        teacher.ID,
			CreatedAt: teacher.CreatedAt,
			UpdatedAt: teacher.UpdatedAt,
			Username:  teacher.Username,
			Email:     teacher.Email,
			Fullname:  teacher.Fullname,
		}
		results = append(results, teacherResponse)
	}
	return results
}
