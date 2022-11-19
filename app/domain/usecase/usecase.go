package usecase

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase/group"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase/interview"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase/mentor"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase/person"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase/course"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase/student"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type Usecases struct {
	Course              CourseUsecaseInterface
	Lecture             LectureUsecaseInterface
	Homework            HomeworkUsecaseInterface
	CertificateTemplate CertificateTemplateUsecaseInterface
	StudentHomework     StudentHomeworkUsecaseInterface
	StudentCertificate  StudentCertificateUsecaseInterface
	CourseLecture       CourseLectureUseCaseInterface
	CourseStatus        CourseStatusUseCaseInterface
	Person              PersonUseCaseInterface
	Student             StudentUseCaseInterface
	StudentGroup        StudentGroupUseCaseInterface
	StudentNote         StudentUseCaseInterface
	Mentor              MentorUseCaseInterface
	GroupContact        GroupContactUseCaseInterface
	Group               GroupUsecaseInterface
	Interview           InterviewUsecaseInterface
	StudentNoteType     StudentNoteUseCaseInterface
	MentorNote          MentorNoteUseCaseInterface
}

func InitUsecases(lg logger.Logger, repo repository.RepositoryInterface) *Usecases {
	return &Usecases{
		Course:              course.NewCourse(lg, repo),
		Lecture:             course.NewLecture(lg, repo),
		Homework:            course.NewHomework(lg, repo),
		CertificateTemplate: course.NewCertificateTemplate(lg, repo),
		StudentHomework:     student.NewStudentHomework(lg, repo),
		StudentCertificate:  student.NewStudentCertificate(lg, repo),
		CourseLecture:       course.NewCourseLecture(lg, repo),
		CourseStatus:        course.NewCourseStatus(lg, repo),
		Person:              person.NewPerson(lg, repo),
		Student:             student.NewStudent(lg, repo),
		StudentGroup:        group.NewStudentGroup(lg, repo),
		StudentNote:         student.NewStudentNote(lg, repo),
		Mentor:              mentor.NewMentor(lg, repo),
		GroupContact:        group.NewGroupContact(lg, repo),
		Group:               group.NewGroup(lg, repo),
		Interview:           interview.NewInterviewUsecase(lg, repo),
	}
}

type ContactUseCaseInterface interface {
}

type ContactTypeUseCaseInterface interface {
}

type CourseUsecaseInterface interface {

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

type CertificateTemplateUsecaseInterface interface {
	GetCertificateTemplates(ctx context.Context) ([]model.CertificateTemplate, error)
	GetCertificateTemplateById(ctx context.Context, id int64) (*model.CertificateTemplate, error)
	AddCertificateTemplate(ctx context.Context, data *model.CertificateTemplate) error
	UpdateCertificateTemplate(ctx context.Context, id int64, data *model.CertificateTemplate) error
	DeleteCertificateTemplate(ctx context.Context, id int64) error
}

type LectureUsecaseInterface interface {
	GetLectures(ctx context.Context) ([]model.Lecture, error)
	GetLectureById(ctx context.Context, id int64) (*model.Lecture, error)
	AddLecture(ctx context.Context, data *model.Lecture) error
	UpdateLecture(ctx context.Context, id int64, data *model.Lecture) error
	DeleteLecture(ctx context.Context, id int64) error
}

type HomeworkUsecaseInterface interface {
	GetHomeworks(ctx context.Context) ([]model.Homework, error)
	GetHomeworksByLectureId(ctx context.Context, lectureId int64) ([]model.Homework, error)
	GetHomeworkById(ctx context.Context, id int64) (*model.Homework, error)
	AddHomework(ctx context.Context, data *model.Homework) error
	UpdateHomework(ctx context.Context, id int64, data *model.Homework) error
	DeleteHomework(ctx context.Context, id int64) error
}

type StudentHomeworkUsecaseInterface interface {
	GetStudentHomeworks(ctx context.Context) ([]model.StudentHomework, error)
	GetStudentHomeworksByStudentId(ctx context.Context, studentId int64) ([]model.StudentHomework, error)
	GetStudentHomeworkById(ctx context.Context, id int64) (*model.StudentHomework, error)
	AddStudentHomework(ctx context.Context, data *model.StudentHomework) error
	UpdateStudentHomework(ctx context.Context, id int64, data *model.StudentHomework) error
	DeleteStudentHomework(ctx context.Context, id int64) error
}

type StudentCertificateUsecaseInterface interface {
	GetStudentCertificates(ctx context.Context) ([]model.StudentCertificate, error)
	GetStudentCertificatesByStudentId(ctx context.Context, studentId int64) ([]model.StudentCertificate, error)
	GetStudentCertificatesByCourseId(ctx context.Context, courseId int64) ([]model.StudentCertificate, error)
	GetStudentCertificateById(ctx context.Context, id int64) (*model.StudentCertificate, error)
	AddStudentCertificate(ctx context.Context, data *model.StudentCertificate) error
	UpdateStudentCertificate(ctx context.Context, id int64, data *model.StudentCertificate) error
	DeleteStudentCertificate(ctx context.Context, id int64) error
}

type MentorNoteUseCaseInterface interface{}
type StudentNoteUseCaseInterface interface{}
type CourseStatusUseCaseInterface interface{}
type StudentGroupUseCaseInterface interface{}
type CourseLectureUseCaseInterface interface{}
type GroupContactUseCaseInterface interface{}

type RoleUsecaseInterface interface {
	GetRoleById(ctx context.Context, id int64) (*model.Role, error)
	GetRoles(ctx context.Context) ([]model.Role, error)
	AddRole(ctx context.Context, role *model.RoleCU) error
	UpadateRoleById(ctx context.Context, role *model.RoleCU) error
	PutRoleById(ctx context.Context, role *model.Role) error
	DeleteRoleById(ctx context.Context, id int64) error
}

type GroupUsecaseInterface interface {
	GetGroupById(ctx context.Context, id int64) (*model.Group, error)
	GetGroups(ctx context.Context) ([]model.Group, error)
	CreateGroup(ctx context.Context, group *model.Group) (int64, error)
	UpdateGroupById(ctx context.Context, id int64, group *model.GroupCU) error
	PutGroupById(ctx context.Context, id int64, group *model.GroupCU) error
	DeleteGroupById(ctx context.Context, id int64) error
}

type InterviewUsecaseInterface interface {
	GetInterviewById(context.Context, int64) (*model.Interview, error)
	GetInterviews(context.Context) ([]model.Interview, error)
	AddInterview(context.Context, *model.Interview) (int64, error)
	UpdateInterviewById(context.Context, int64, *model.Interview) error
	DeleteInterviewById(context.Context, int64) error
}
