package v1

import (
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type Handlers struct {
	Course        *CourseHandlers
	CourseLecture *CourseLectureHandlers
	CourseStatus  *CourseStatusHandlers

	Person *PersonHandlers

	Student         *StudentHandlers
	StudentGroup    *StudentGroupHandlers
	StudentNote     *StudentNoteHandlers
	StudentNoteType *StudentNoteTypeHandlers

	Mentor     *MentorHandlers
	MentorNote *MentorNoteHandlers

	GroupContact *GroupContactHandlers
}

func InitHandlers(lg logger.Logger, u usecase.Usecases) *Handlers {
	return &Handlers{
		Course:        NewCourseHandler(lg, u.Course),
		CourseLecture: NewCourseLectureHandler(lg, u.CourseLecture),
		CourseStatus:  NewCourseStatusHandler(lg, u.CourseStatus),
		Person:        NewPersonHandler(lg, u.Person),
		Student:       NewStudentHandler(lg, u.Student),
		StudentGroup:  NewStudentGroupHandler(lg, u.StudentGroup),
		StudentNote:   NewStudentNoteHandler(lg, u.StudentNote),
		Mentor:        NewMentorHandler(lg, u.Mentor),
	}
}
