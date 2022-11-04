package v1

import (
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type Handlers struct {
	Lecture             *LectureHandlers
	Homework            *HomeworkHandlers
	Project             *ProjectHandlers
	CertificateTemplate *CertificateTemplateHandlers
	StudentHomework     *StudentHomeworkHandlers
	StudentCertificate  *StudentCertificateHandlers
	LogOperation        *LogOperationHandlers
	Operation           *OperationHandlers
	OperationType       *OperationTypeHandlers
	ContactType         *ContactTypeHandlers
	Course              *CourseHandlers
	CourseLecture       *CourseLectureHandlers
	CourseStatus        *CourseStatusHandlers
	Person              *PersonHandlers
	Student             *StudentHandlers
	StudentGroup        *StudentGroupHandlers
	StudentNote         *StudentNoteHandlers
	StudentNoteType     *StudentNoteTypeHandlers
	Mentor              *MentorHandlers
	MentorNote          *MentorNoteHandlers
	GroupContact        *GroupContactHandlers
}

func InitHandlers(lg logger.Logger, u usecase.Usecases) *Handlers {
	return &Handlers{

		CourseLecture:       NewCourseLectureHandler(lg, u.CourseLecture),
		CourseStatus:        NewCourseStatusHandler(lg, u.CourseStatus),
		Person:              NewPersonHandler(lg, u.Person),
		Student:             NewStudentHandler(lg, u.Student),
		StudentGroup:        NewStudentGroupHandler(lg, u.StudentGroup),
		StudentNote:         NewStudentNoteHandler(lg, u.StudentNote),
		Mentor:              NewMentorHandler(lg, u.Mentor),
		Course:              NewCourseHandler(lg, u.Course),
		Lecture:             NewLectureHandler(lg, u.Lecture),
		Homework:            NewHomeworkHandler(lg, u.Homework),
		Project:             NewProjectHandler(lg, u.Project),
		CertificateTemplate: NewCertificateTemplateHandler(lg, u.CertificateTemplate),
		StudentHomework:     NewStudentHomeworkHandler(lg, u.StudentHomework),
		StudentCertificate:  NewStudentCertificateHandler(lg, u.StudentCertificate),
		LogOperation:        NewLogOperationHandler(lg, u.LogOperation),
		Operation:           NewOperationHandler(lg, u.Operation),
		OperationType:       NewOperationTypeHandler(lg, u.OperationType),
		ContactType:         NewContactTypeHandler(lg, u.ContactType),
	}
}
