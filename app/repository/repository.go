package repository

import (
	"github.com/lozovoya/GolangUnitedSchool/app/domain/model"
	"golang.org/x/net/context"
)

type RepositoryInterface interface {
	GetPersons(ctx context.Context) ([]model.Person, error)
	GetPersonById(ctx context.Context, id int64) (*model.Person, error)
	AddNewPerson(ctx context.Context, data *model.NewPersonDto) (int64, error)
	UpdatePersonByID(ctx context.Context, id int64, data *model.UpdatePerson) error
	PutPersonByID(ctx context.Context, id int64, data *model.UpdatePerson) error
	DeletePersonById(ctx context.Context, id int64) error
	//student //
	GetStudents(ctx context.Context) ([]model.Student, error)
	GetStudentByStudentId(ctx context.Context, id int64) (*model.Student, error)
	AddStudent(ctx context.Context, data *model.Student) (int64, error)
	UpdateStudentByStudentId(ctx context.Context, id int64, data *model.StudentUpdate) error
	DeleteStudentByStudentId(ctx context.Context, id int64) error
	PutStudentByStudentId(ctx context.Context, id int64, data *model.StudentUpdate) error
	//mentor
	GetMentors(ctx context.Context) ([]model.Mentor, error)
	GetMentorById(ctx context.Context, id int64) (*model.Mentor, error)
	AddMentor(ctx context.Context, data *model.UpdateMentor) (int64, error)
	UpdateMentorByMentorId(ctx context.Context, id int64, data *model.UpdateMentor) error
	DeleteMentorByMentorId(ctx context.Context, id int64) error
	PutMentorById(ctx context.Context, id int64, data *model.UpdateMentor) error
	//mentor note
	GetMentorNotes(ctx context.Context) ([]model.MentorNote, error)
	GetMentorNotesByMentorId(ctx context.Context, id int64) ([]model.MentorNote, error)
	GetMentorNoteByMentorNoteId(ctx context.Context, id int64) (*model.MentorNote, error)
	AddMentorNote(ctx context.Context, data *model.NewMentorNote) (int64, error)
	UpdateMentorNoteByMentorNoteId(ctx context.Context, id int64, data *model.UpdateMentorNote) error
	DeleteMentorNoteByMentorNoteId(ctx context.Context, id int64) error
	PutMentorNoteByMentorNoteId(ctx context.Context, id int64, data *model.UpdateMentorNote) error
	//student note
	GetStudentNotes(ctx context.Context) ([]model.StudentNote, error)
	GetStudentNoteById(ctx context.Context, id int64) (*model.StudentNote, error)
	AddStudentNote(ctx context.Context, data *model.NewStudentNote) (int64, error)
	UpdateStudentNoteByStudentId(ctx context.Context, id int64, data *model.UpdateStudentNote) error
	PutStudentNoteById(ctx context.Context, id int64, data *model.UpdateStudentNote) error
	DeleteStudentNoteByStudentNoteId(ctx context.Context, id int64) error
	GetStudentsNotesByStudentId(ctx context.Context, id int64) ([]model.StudentNote, error)

	//student note type
	GetStudentNoteTypes(ctx context.Context) ([]model.StudentNoteType, error)
	GetStudentNoteTypeById(ctx context.Context, id int64) (*model.StudentNoteType, error)
	AddStudentNoteType(ctx context.Context, data *model.NewStudentNoteType) (int64, error)
	UpdateStudentNoteTypeById(ctx context.Context, id int64, data *model.UpdateStudentNoteType) error
	PutStudentNoteTypeById(ctx context.Context, id int64, data *model.UpdateStudentNoteType) error
	DeleteStudentNoteTypeById(ctx context.Context, id int64) error
	//group contact
	GetAllGroupContacts(ctx context.Context) ([]model.GroupContact, error)
	GetGroupContactById(ctx context.Context, id int64) (*model.GroupContact, error)
	AddGroupContact(ctx context.Context, data *model.GroupContactCU) (int64, error)
	PutGroupContactById(ctx context.Context, id int64, data *model.GroupContactCU) error
	UpdateGroupContactById(ctx context.Context, id int64, data *model.GroupContactUpdate) error
	DeleteGroupContactById(ctx context.Context, id int64) error
	GetGroupContacts(ctx context.Context, id int64) ([]model.GroupContact, error)

	// student group
	GetStudentGroups(ctx context.Context) ([]model.StudentGroup, error)
	GetStudentGroupById(ctx context.Context, id int64) (*model.StudentGroup, error)
	CreateStudentGroup(ctx context.Context, data *model.StudentGroup) (int64, error)
	PutStudentGroupById(ctx context.Context, id int64, data *model.UpdateStudentGroup) error
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
	UpdateRoleById(ctx context.Context, id int64, role *model.RoleCU) error
	PutRoleById(ctx context.Context, id int64, role *model.Role) error
	DeleteRoleById(ctx context.Context, id int64) error

	//// group
	//GetGroupById(ctx context.Context, id int64) (*model.Group, error)
	//GetGroups(ctx context.Context) ([]model.Group, error)
	//CreateGroup(ctx context.Context, group *model.Group) (int64, error)
	//UpdateGroupById(ctx context.Context, id int64, group *model.GroupCU) error
	//PutGroupById(ctx context.Context, id int64, group *model.GroupCU) error
	//DeleteGroupById(ctx context.Context, id int64) error

	GetProjects(ctx context.Context) ([]model.Project, error)
	GetProjectsByCourseId(ctx context.Context, courseId int64) ([]model.Project, error)
	GetProjectsByGroupId(ctx context.Context, groupId int64) ([]model.Project, error)
	GetProjectById(ctx context.Context, id int64) (*model.Project, error)
	CreateProject(ctx context.Context, data *model.Project) error
	UpdateProject(ctx context.Context, id int64, data *model.Project) error
	DeleteProject(ctx context.Context, id int64) error
}
