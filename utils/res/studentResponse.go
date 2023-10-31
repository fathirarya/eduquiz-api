package res

import (
	"eduquiz-api/model/domain"
	"eduquiz-api/model/schema"
	"eduquiz-api/model/web"
)

func StudentSchemaToStudentDomain(student *schema.Student) *domain.Student {
	return &domain.Student{
		ID:       student.ID,
		Username: student.Username,
		Password: student.Password,
		Email:    student.Email,
		Fullname: student.Fullname,
	}
}

func StudentDomainToStudentResponse(student *domain.Student) web.StudentResponse {
	return web.StudentResponse{
		ID:        student.ID,
		CreatedAt: student.CreatedAt,
		UpdatedAt: student.UpdatedAt,
		Username:  student.Username,
		Email:     student.Email,
		Fullname:  student.Fullname,
	}
}

func StudentDomainToStudentLoginResponse(student *domain.Student) web.StudentLoginResponse {
	return web.StudentLoginResponse{
		Fullname: student.Fullname,
		Email:    student.Email,
	}
}

func ConvertStudentResponse(students []domain.Student) []web.StudentResponse {
	var results []web.StudentResponse
	for _, student := range students {
		studentResponse := web.StudentResponse{
			ID:        student.ID,
			CreatedAt: student.CreatedAt,
			UpdatedAt: student.UpdatedAt,
			Username:  student.Username,
			Email:     student.Email,
			Fullname:  student.Fullname,
		}
		results = append(results, studentResponse)
	}
	return results
}
