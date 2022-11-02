package repository

import (
	m "github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"golang.org/x/net/context"
)

type cntxt context.Context

type RepositoryInterface interface {
	//Person1
	GetPersons(ctx cntxt) ([]m.Person, error)
	GetPersonById(ctx cntxt, id int64) (*m.Person, error)
	AddNewPerson(ctx cntxt, data *m.Person) error
	EditPersonById(ctx cntxt, id int64, data *m.Person) error
	DeletePersonById(ctx cntxt, id int64) error
	//student1
	GetStudents(ctx cntxt) ([]m.Student, error)
	GetStudentByStudentId(ctx cntxt, id int64) (*m.Student, error)
	AddStudent(ctx cntxt, data *m.Student) error
	EditStudentByStudentId(ctx cntxt, id int64, data *m.Student) error
	DeleteStudentByStudentId(ctx cntxt, id int64) error
	//mentor1
	GetMentors(ctx cntxt) ([]m.Mentor, error)
	GetMentorById(ctx cntxt, id int64) (*m.Mentor, error)
	AddMentor(ctx cntxt, data *m.Mentor) error
	EditMentorByMentorId(ctx cntxt, id int64, data *m.Mentor) error
	DeleteMentorByMentorId(ctx cntxt, id int64) error
	//mentor note1
	GetMentorNotes(ctx cntxt) ([]m.MentorNote, error)
	GetMentorNotesByMentorId(ctx cntxt, id int64) ([]m.MentorNote, error)
	GetMentorNoteByMentorNoteId(ctx cntxt, id int64) (*m.MentorNote, error)
	AddMentorNote(ctx cntxt, data *m.MentorNote) error
	EditMentorNoteByMentorNoteId(ctx cntxt, id int64, data *m.MentorNote) error
	DeleteMentorNoteByMentorNoteId(ctx cntxt, id int64) error
	//student note1
	GetStudentNotes(ctx cntxt) ([]m.StudentNote, error)
	GetStudentNoteByStudentId(ctx cntxt, id int64) (*m.StudentNote, error)
	AddStudentNote(ctx cntxt, data *m.StudentNote) error
	EditStudentNoteByStudentId(ctx cntxt, id int64, data *m.StudentNote) error
	DeleteStudentNoteByStudentNoteId(ctx cntxt, id int64) error
	//student note type1
	GetStudentNoteTypes(ctx cntxt) ([]m.StudentNoteType, error)
	GetStudentNoteTypeById(ctx cntxt, id int64) (*m.StudentNoteType, error)
	AddStudentNoteType(ctx cntxt, data *m.StudentNoteType) error
	EditStudentNoteTypeById(ctx cntxt, id int64, data *m.StudentNoteType) error
	DeleteStudentNoteTypeById(ctx cntxt, is int64) error
	//group contact1
	GetGroupContacts(ctx cntxt) ([]m.GroupContact, error)
	GetGroupContactById(ctx cntxt, id int64) (*m.GroupContact, error)
	AddGroupContact(ctx, data *m.GroupContact) error
	EditGroupContactById(ctx cntxt, id int64, data *m.GroupContact) error
	DeleteGroupContactById(ctx, id int64) error
	// student group1
	GetStudentGroups(ctx cntxt) ([]m.StudentGroup, error)
	GetStudentGroupById(ctx cntxt, id int64) (*m.StudentGroup, error)
	AddStudentGroup(ctx cntxt, data *m.StudentGroup) error
	EditStudentGroupById(ctx cntxt, id int64, data *m.StudentGroup) error
	DeleteStudentGroupById(ctx cntxt, id int64) error
	// course status1
	GetCourseStatuses(ctx cntxt) ([]m.CourseStatus, error)
	GetCourseStatusById(ctx cntxt, id int64) (*m.CourseStatus, error)
	AddCourseStatus(ctx cntxt, data *m.CourseStatus) error
	EditCourseStatusById(ctx cntxt, id int64, data *m.CourseStatus) error
	DeleteCourseStatusById(ctx cntxt, id int64) error
	// course lecture
	GetCourseLectures(ctx cntxt) ([]m.CourseLecture, error)
	GetCourseLectureById(ctx cntxt, id int64) (*m.CourseLecture, error)
	AddCourseLecture(ctx cntxt, data *m.CourseLecture) error
	EditCourseLectureById(ctx cntxt, id int64, data *m.CourseLecture) error
	DeleteCourseLectureById(ctx cntxt, id int64) error
}
