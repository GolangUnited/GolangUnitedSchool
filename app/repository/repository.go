package repository

import (
	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"golang.org/x/net/context"
)

type RepositoryInterface interface {
	//Person1
	GetPersons(ctx context.Context) ([]model.Person, error)
	GetPersonById(ctx context.Context, id int64) (*model.Person, error)
	AddNewPerson(ctx context.Context, data *model.Person) error
	EditPersonById(ctx context.Context, id int64, data *model.Person) error
	DeletePersonById(ctx context.Context, id int64) error
	//student1
	GetStudents(ctx context.Context) ([]model.Student, error)
	GetStudentByStudentId(ctx context.Context, id int64) (*model.Student, error)
	AddStudent(ctx context.Context, data *model.Student) error
	EditStudentByStudentId(ctx context.Context, id int64, data *model.Student) error
	DeleteStudentByStudentId(ctx context.Context, id int64) error
	//mentor1
	GetMentors(ctx context.Context) ([]model.Mentor, error)
	GetMentorById(ctx context.Context, id int64) (*model.Mentor, error)
	AddMentor(ctx context.Context, data *model.Mentor) error
	EditMentorByMentorId(ctx context.Context, id int64, data *model.Mentor) error
	DeleteMentorByMentorId(ctx context.Context, id int64) error
	//mentor note1
	GetMentorNotes(ctx context.Context) ([]model.MentorNote, error)
	GetMentorNotesByMentorId(ctx context.Context, id int64) ([]model.MentorNote, error)
	GetMentorNoteByMentorNoteId(ctx context.Context, id int64) (*model.MentorNote, error)
	AddMentorNote(ctx context.Context, data *model.MentorNote) error
	EditMentorNoteByMentorNoteId(ctx context.Context, id int64, data *model.MentorNote) error
	DeleteMentorNoteByMentorNoteId(ctx context.Context, id int64) error
	//student note1
	GetStudentNotes(ctx context.Context) ([]model.StudentNote, error)
	GetStudentNoteByStudentId(ctx context.Context, id int64) (*model.StudentNote, error)
	AddStudentNote(ctx context.Context, data *model.StudentNote) error
	EditStudentNoteByStudentId(ctx context.Context, id int64, data *model.StudentNote) error
	DeleteStudentNoteByStudentNoteId(ctx context.Context, id int64) error
	//student note type1
	GetStudentNoteTypes(ctx context.Context) ([]model.StudentNoteType, error)
	GetStudentNoteTypeById(ctx context.Context, id int64) (*model.StudentNoteType, error)
	AddStudentNoteType(ctx context.Context, data *model.StudentNoteType) error
	EditStudentNoteTypeById(ctx context.Context, id int64, data *model.StudentNoteType) error
	DeleteStudentNoteTypeById(ctx context.Context, is int64) error
	//group contact1
	GetGroupContacts(ctx context.Context) ([]model.GroupContact, error)
	GetGroupContactById(ctx context.Context, id int64) (*model.GroupContact, error)
	AddGroupContact(ctx, data *model.GroupContact) error
	EditGroupContactById(ctx context.Context, id int64, data *model.GroupContact) error
	DeleteGroupContactById(ctx, id int64) error
	// student group1
	GetStudentGroups(ctx context.Context) ([]model.StudentGroup, error)
	GetStudentGroupById(ctx context.Context, id int64) (*model.StudentGroup, error)
	AddStudentGroup(ctx context.Context, data *model.StudentGroup) error
	EditStudentGroupById(ctx context.Context, id int64, data *model.StudentGroup) error
	DeleteStudentGroupById(ctx context.Context, id int64) error
	// course status1
	GetCourseStatuses(ctx context.Context) ([]model.CourseStatus, error)
	GetCourseStatusById(ctx context.Context, id int64) (*model.CourseStatus, error)
	AddCourseStatus(ctx context.Context, data *model.CourseStatus) error
	EditCourseStatusById(ctx context.Context, id int64, data *model.CourseStatus) error
	DeleteCourseStatusById(ctx context.Context, id int64) error
	// course lecture
	GetCourseLectures(ctx context.Context) ([]model.CourseLecture, error)
	GetCourseLectureById(ctx context.Context, id int64) (*model.CourseLecture, error)
	AddCourseLecture(ctx context.Context, data *model.CourseLecture) error
	EditCourseLectureById(ctx context.Context, id int64, data *model.CourseLecture) error
	DeleteCourseLectureById(ctx context.Context, id int64) error

	GetCertificateTemplates(ctx context.Context) ([]model.CertificateTemplate, error)
	GetCertificateTemplateById(ctx context.Context, id int64) (*model.CertificateTemplate, error)
	AddCertificateTemplate(ctx context.Context, data *model.CertificateTemplate) error
	UpdateCertificateTemplate(ctx context.Context, id int64, data *model.CertificateTemplate) error
	DeleteCertificateTemplate(ctx context.Context, id int64) error

	GetLectures(ctx context.Context) ([]model.Lecture, error)
	GetLectureById(ctx context.Context, id int64) (*model.Lecture, error)
	AddLecture(ctx context.Context, data *model.Lecture) error
	UpdateLecture(ctx context.Context, id int64, data *model.Lecture) error
	DeleteLecture(ctx context.Context, id int64) error

	GetHomeworks(ctx context.Context) ([]model.Homework, error)
	GetHomeworksByLectureId(ctx context.Context, lectureId int64) ([]model.Homework, error)
	GetHomeworkById(ctx context.Context, id int64) (*model.Homework, error)
	AddHomework(ctx context.Context, data *model.Homework) error
	UpdateHomework(ctx context.Context, id int64, data *model.Homework) error
	DeleteHomework(ctx context.Context, id int64) error

	GetStudentHomeworks(ctx context.Context) ([]model.StudentHomework, error)
	GetStudentHomeworksByStudentId(ctx context.Context, studentId int64) ([]model.StudentHomework, error)
	GetStudentHomeworkById(ctx context.Context, id int64) (*model.StudentHomework, error)
	AddStudentHomework(ctx context.Context, data *model.StudentHomework) error
	UpdateStudentHomework(ctx context.Context, id int64, data *model.StudentHomework) error
	DeleteStudentHomework(ctx context.Context, id int64) error

	GetStudentCertificates(ctx context.Context) ([]model.StudentCertificate, error)
	GetStudentCertificatesByStudentId(ctx context.Context, studentId int64) ([]model.StudentCertificate, error)
	GetStudentCertificatesByCourseId(ctx context.Context, courseId int64) ([]model.StudentCertificate, error)
	GetStudentCertificateById(ctx context.Context, id int64) (*model.StudentCertificate, error)
	AddStudentCertificate(ctx context.Context, data *model.StudentCertificate) error
	UpdateStudentCertificate(ctx context.Context, id int64, data *model.StudentCertificate) error
	DeleteStudentCertificate(ctx context.Context, id int64) error
}
