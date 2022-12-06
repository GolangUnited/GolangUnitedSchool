package repository

import (
	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"golang.org/x/net/context"
)

type RepositoryInterface interface {
	//Person //
	GetPersons(ctx context.Context) (model.PersonListDto, error)
	GetPersonById(ctx context.Context, id int64) (*model.Person, error)
	AddNewPerson(ctx context.Context, data *model.NewPersonDto) error
	UpdatePersonById(ctx context.Context, id int64, data model.UpdatePerson) error
	PutPersonById(ctx context.Context, id int64, data *model.Person) error
	DeletePersonById(ctx context.Context, id int64) error
	//student //
	GetStudents(ctx context.Context) (model.StudentsListDto, error)
	GetStudentByStudentId(ctx context.Context, id int64) (*model.Student, error)
	AddStudent(ctx context.Context, data *model.Student) error
	PutStudentByStudentId(ctx context.Context, id int64, data *model.Student) error
	DeleteStudentByStudentId(ctx context.Context, id int64) error
	//mentor
	GetMentors(ctx context.Context) ([]model.Mentor, error)
	GetMentorById(ctx context.Context, id int64) (*model.Mentor, error)
	AddMentor(ctx context.Context, data *model.Mentor) error
	UpdateMentorByMentorId(ctx context.Context, id int64, data *model.Mentor) error
	DeleteMentorByMentorId(ctx context.Context, id int64) error
	//mentor note
	GetMentorNotes(ctx context.Context) ([]model.MentorNote, error)
	GetMentorNotesByMentorId(ctx context.Context, id int64) ([]model.MentorNote, error)
	GetMentorNoteByMentorNoteId(ctx context.Context, id int64) (*model.MentorNote, error)
	AddMentorNote(ctx context.Context, data *model.MentorNote) error
	UpdateMentorNoteByMentorNoteId(ctx context.Context, id int64, data *model.MentorNote) error
	DeleteMentorNoteByMentorNoteId(ctx context.Context, id int64) error
	//student note
	GetStudentNotes(ctx context.Context) ([]model.StudentNote, error)
	GetStudentNoteByStudentId(ctx context.Context, id int64) (*model.StudentNote, error)
	AddStudentNote(ctx context.Context, data *model.StudentNote) error
	UpdateStudentNoteByStudentId(ctx context.Context, id int64, data *model.StudentNote) error
	DeleteStudentNoteByStudentNoteId(ctx context.Context, id int64) error
	//student note type
	GetStudentNoteTypes(ctx context.Context) ([]model.StudentNoteType, error)
	GetStudentNoteTypeById(ctx context.Context, id int64) (*model.StudentNoteType, error)
	AddStudentNoteType(ctx context.Context, data *model.StudentNoteType) error
	UpdateStudentNoteTypeById(ctx context.Context, id int64, data *model.StudentNoteType) error
	DeleteStudentNoteTypeById(ctx context.Context, is int64) error
	//group contact
	GetGroupContacts(ctx context.Context) ([]model.GroupContact, error)
	GetGroupContactById(ctx context.Context, id int64) (*model.GroupContact, error)
	AddGroupContact(ctx context.Context, data *model.GroupContact) error
	UpdateGroupContactById(ctx context.Context, id int64, data *model.GroupContact) error
	DeleteGroupContactById(ctx context.Context, id int64) error
	// student group
	GetStudentGroups(ctx context.Context) ([]model.StudentGroup, error)
	GetStudentGroupById(ctx context.Context, id int64) (*model.StudentGroup, error)
	AddStudentGroup(ctx context.Context, data *model.StudentGroup) error
	UpdateStudentGroupById(ctx context.Context, id int64, data *model.StudentGroup) error
	DeleteStudentGroupById(ctx context.Context, id int64) error
	// course status
	GetCourseStatuses(ctx context.Context) ([]model.CourseStatus, error)
	GetCourseStatusById(ctx context.Context, id int64) (*model.CourseStatus, error)
	AddCourseStatus(ctx context.Context, data *model.CourseStatus) (int64, error)
	UpdateCourseStatusById(ctx context.Context, id int64, data *model.CourseStatus) error
	DeleteCourseStatusById(ctx context.Context, id int64) error
	// course lecture
	// GetCourseLectures(ctx context.Context) ([]model.CourseLecture, error)
	// GetCourseLectureById(ctx context.Context, id int64) (*model.CourseLecture, error)
	// AddCourseLecture(ctx context.Context, data *model.CourseLecture) error
	// UpdateCourseLectureById(ctx context.Context, id int64, data *model.CourseLecture) error
	// DeleteCourseLectureById(ctx context.Context, id int64) error

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
	// CreateCourse(context.Context, *model.CourseCreate) (int64, error)
	// GetCourseByID(context.Context, int64) (*model.Course, error)
	// GetCourses(context.Context, *model.PaginationParams) (*model.CourseList, error)
	// UpdateCourseByID(context.Context, int64, *model.CourseUpdate) error
	// PutCourseByID(context.Context, int64, *model.CourseUpdate) error
	// DeleteCourseByID(context.Context, int64) error

	// interview
	GetInterviewById(context.Context, int64) (*model.Interview, error)
	GetInterviews(context.Context) ([]model.Interview, error)
	AddInterview(context.Context, *model.Interview) (int64, error)
	UpdateInterviewById(context.Context, int64, *model.Interview) error
	DeleteInterviewById(context.Context, int64) error

	// role
	GetRoleById(ctx context.Context, id int64) (*model.Role, error)
	GetRoles(ctx context.Context) ([]model.Role, error)
	AddRole(ctx context.Context, role *model.RoleCU) (int64, error)
	UpadateRoleById(ctx context.Context, id int64, role *model.RoleCU) error
	PutRoleById(ctx context.Context, id int64, role *model.Role) error
	DeleteRoleById(ctx context.Context, id int64) error

	// group
	GetGroupById(ctx context.Context, id int64) (*model.Group, error)
	GetGroups(ctx context.Context) ([]model.Group, error)
	CreateGroup(ctx context.Context, group *model.Group) (int64, error)
	UpdateGroupById(ctx context.Context, id int64, group *model.GroupCU) error
	PutGroupById(ctx context.Context, id int64, group *model.GroupCU) error
	DeleteGroupById(ctx context.Context, id int64) error

	GetProjects(ctx context.Context) ([]model.Project, error)
	GetProjectsByCourseId(ctx context.Context, courseId int64) ([]model.Project, error)
	GetProjectsByGroupId(ctx context.Context, groupId int64) ([]model.Project, error)
	GetProjectById(ctx context.Context, id int64) (*model.Project, error)
	CreateProject(ctx context.Context, data *model.Project) error
	UpdateProject(ctx context.Context, id int64, data *model.Project) error
	DeleteProject(ctx context.Context, id int64) error
}
