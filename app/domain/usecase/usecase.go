package usecase

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase/group"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase/mentor"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase/person"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase/contact"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase/course"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase/operation"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase/student"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
	"github.com/lozovoya/GolangUnitedSchool/app/repository"
)

type Usecases struct {
	Course              CourseUsecaseInterface
	Lecture             LectureUsecaseInterface
	Homework            HomeworkUsecaseInterface
	Project             ProjectUseCaseInterface
	CertificateTemplate CertificateTemplateUsecaseInterface
	StudentHomework     StudentHomeworkUsecaseInterface
	StudentCertificate  StudentCertificateUsecaseInterface
	LogOperation        LogOperationUsecaseInterface
	Operation           OperationUsecaseInterface
	OperationType       OperationTypeUsecaseInterface
	ContactType         ContactTypeUsecaseInterface
	CourseLecture       CourseLectureUseCaseInterface
	CourseStatus        CourseStatusUseCaseInterface
	Person              PersonUseCaseInterface
	Student             StudentUseCaseInterface
	StudentGroup        StudentGroupUseCaseInterface
	StudentNote         StudentUseCaseInterface
	Mentor              MentorUseCaseInterface
	GroupContact        GroupContactUseCaseInterface
}

func InitUsecases(lg logger.Logger, repo repository.RepositoryInterface) *Usecases {
	return &Usecases{
		Course:              course.NewCourse(lg, repo),
		Lecture:             course.NewLecture(lg, repo),
		Homework:            course.NewHomework(lg, repo),
		Project:             course.NewProject(lg, repo),
		CertificateTemplate: course.NewCertificateTemplate(lg, repo),
		StudentHomework:     student.NewStudentHomework(lg, repo),
		StudentCertificate:  student.NewStudentCertificate(lg, repo),
		LogOperation:        operation.NewLogOperation(lg, repo),
		Operation:           operation.NewOperation(lg, repo),
		OperationType:       operation.NewOperationType(lg, repo),
		ContactType:         contact.NewContactType(lg, repo),
		CourseLecture:       course.NewCourseLecture(lg, repo),
		CourseStatus:        course.NewCourseStatus(lg, repo),
		Person:              person.NewPerson(lg, repo),
		Student:             student.NewStudent(lg, repo),
		StudentGroup:        group.NewStudentGroup(lg, repo),
		StudentNote:         student.NewStudentNote(lg, repo),
		Mentor:              mentor.NewMentor(lg, repo),
		GroupContact:        group.NewGroupContact(lg, repo),
	}
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
type ProjectUseCaseInterface interface {
	GetProjects(ctx context.Context) ([]model.Project, error)
	GetProjectById(ctx context.Context, id int64) (*model.Project, error)
	AddProject(ctx context.Context, data *model.Project) error
	UpdateProject(ctx context.Context, id int64, data *model.Project) error
	DeleteProject(ctx context.Context, id int64) error
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

type LogOperationUsecaseInterface interface {
	GetLogOperationById(ctx context.Context, id int64) (*model.LogOperation, error)
	AddLogOperation(ctx context.Context, data *model.LogOperation) error
	DeleteLogOperation(ctx context.Context, id int64) error
}

type OperationUsecaseInterface interface {
	GetOperations(ctx context.Context) ([]model.Operation, error)
	GetOperationById(ctx context.Context, id int64) (*model.Operation, error)
	AddOperation(ctx context.Context) error
	UpdateOperation(ctx context.Context, id int64) error
	DeleteOperation(ctx context.Context, id int64) error
}

type OperationTypeUsecaseInterface interface {
	GetOperationTypes(ctx context.Context) ([]model.OperationType, error)
	GetOperationTypeById(ctx context.Context, id int64) (*model.OperationType, error)
	AddOperationType(ctx context.Context) error
	UpdateOperationType(ctx context.Context, id int64) error
	DeleteOperationType(ctx context.Context, id int64) error
}

type ContactTypeUsecaseInterface interface {
	GetContactTypes(ctx context.Context) ([]model.ContactType, error)
	GetContactTypeById(ctx context.Context, id int64) (*model.ContactType, error)
	AddContactType(ctx context.Context) error
	UpdateContactType(ctx context.Context, id int64) error
	DeleteContactType(ctx context.Context, id int64) error
}

type MentorNoteUseCaseInterface interface{}
type StudentNoteUseCaseInterface interface{}
type CourseStatusUseCaseInterface interface{}
type StudentGroupUseCaseInterface interface{}
type CourseLectureUseCaseInterface interface{}
type GroupContactUseCaseInterface interface{}
