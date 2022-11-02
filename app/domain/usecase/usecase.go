package usecase

import (
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type Usecases struct {
	Course        CourseUseCaseInterface
	CourseLecture CourseLectureUseCaseInterface
	CourseStatus  CourseStatusUseCaseInterface
	Person        PersonUseCaseInterface
	Student       StudentUseCaseInterface
	StudentGroup  StudentGroupUseCaseInterface
	StudentNote   StudentUseCaseInterface
	Mentor        MentorUseCaseInterface
	GroupContact  GroupContactUseCaseInterface
}

func InitUsecases(lg logger.Logger, repo repository.RepositoryInterface) *Usecases {
	return &Usecases{
		Course:        NewCourse(lg, repo),
		CourseLecture: NewCourseLecture(lg, repo),
		CourseStatus:  NewCourseStatus(lg, repo),
		Person:        NewPerson(lg, repo),
		Student:       NewStudent(lg, repo),
		StudentGroup:  NewStudentGroup(lg, repo),
		StudentNote:   NewStudentNote(lg, repo),
		Mentor:        NewMentor(lg, repo),
		GroupContact:  NewGroupContact(lg, repo),
	}
}

type CourseUseCaseInterface interface {

	// AddCourse()
	// EditCourse()
	// EditCourseById(id int64) error
	// DeleteCourse() error
	// DeleteCourseById(id int64) error
}

type PersonUseCaseInterface interface {
	// GetPersonByField() (Person, error)
	// GetPersonById(int64) (Person, error)
	// AddNewPerson() error
	// UpdatePersonById (int64) error
	// RemovePersonById (int64) error
}
type StudentUseCaseInterface interface {
	//AddStudent
	//DeleteStudent
	//DeleteStudentByStudentId
}
type MentorUseCaseInterface interface {
}

type MentorNoteUseCaseInterface interface{}
type StudentNoteUseCaseInterface interface{}
type CourseStatusUseCaseInterface interface{}
type StudentGroupUseCaseInterface interface{}
type CourseLectureUseCaseInterface interface{}
type GroupContactUseCaseInterface interface{}
