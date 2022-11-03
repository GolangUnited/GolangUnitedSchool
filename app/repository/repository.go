package repository

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
)

// RepositoryInterface is main interface for all repo implementations
type RepositoryInterface interface {
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

	// Course repo interfaces
	CreateCourse(context.Context, *model.CourseCreate) (int64, error)
	GetCourseByID(context.Context, int64) (*model.Course, error)
	GetCourses(context.Context, *model.PaginationParams) (*model.CourseList, error)
	UpdateCourseByID(context.Context, int64, *model.CourseUpdate) error
	PutCourseByID(context.Context, int64, *model.CourseUpdate) error
	DeleteCourseByID(context.Context, int64) error
}
