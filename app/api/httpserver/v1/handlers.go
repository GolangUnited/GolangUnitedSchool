package v1

import (
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type Handlers struct {
	Course              *CourseHandlers
	Person              *PersonHandlers
	Lecture             *LectureHandlers
	Homework            *HomeworkHandlers
	CertificateTemplate *CertificateTemplateHandlers
	StudentHomework     *StudentHomeworkHandlers
	StudentCertificate  *StudentCertificateHandlers
}

func InitHandlers(lg logger.Logger, u usecase.Usecases) *Handlers {
	return &Handlers{
		Course:              NewCourseHandler(lg, u.Course),
		Person:              NewPersonHandler(lg),
		Lecture:             NewLectureHandler(lg, u.Lecture),
		Homework:            NewHomeworkHandler(lg, u.Homework),
		CertificateTemplate: NewCertificateTemplateHandler(lg, u.CertificateTemplate),
		StudentHomework:     NewStudentHomeworkHandler(lg, u.StudentHomework),
		StudentCertificate:  NewStudentCertificateHandler(lg, u.StudentCertificate),
	}
}
