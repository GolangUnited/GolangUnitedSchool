package v1

import (
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type Handlers struct {
	Lecture             *LectureHandlers
	Homework            *HomeworkHandlers
	CertificateTemplate *CertificateTemplateHandlers
	StudentHomework     *StudentHomeworkHandlers
	StudentCertificate  *StudentCertificateHandlers
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
	Project             *ProjectHandlers
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
		CertificateTemplate: NewCertificateTemplateHandler(lg, u.CertificateTemplate),
		StudentHomework:     NewStudentHomeworkHandler(lg, u.StudentHomework),
		StudentCertificate:  NewStudentCertificateHandler(lg, u.StudentCertificate),
		GroupContact:        NewGroupContactHandler(lg, u.GroupContact),
		StudentNoteType:     NewStudentNoteTypeHandler(lg, u.StudentNoteType),
		MentorNote:          NewMentorNoteHandler(lg, u.MentorNote),
		Project:             NewProjectHandler(lg, u.Project),
	}
}
