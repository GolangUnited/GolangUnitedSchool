package httpserver

import (
	"net/http"

	_ "github.com/lozovoya/GolangUnitedSchool/docs"

	"github.com/gin-gonic/gin"
	v1 "github.com/lozovoya/GolangUnitedSchool/app/api/httpserver/v1"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(
	handlers *v1.Handlers,
) *gin.Engine {
	router := gin.Default()

	// hello page
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello everything is ok!")
	})

	// docs route
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/v1")
	courseRouter(api, handlers)
	personRouter(api, handlers)
	studentRouter(api, handlers)
	mentorRouter(api, handlers)
	mentorNoteRouter(api, handlers)
	studentNoteRouter(api, handlers)
	studentNoteTypeRouter(api, handlers)
	groupContactRouter(api, handlers)
	studentGroupRouter(api, handlers)
	courseStatusRouter(api, handlers)
	courseLectureRouter(api, handlers)
	lectureRouter(api, handlers.Lecture)
	homeworkRouter(api, handlers.Homework)
	certificateTemplateRouter(api, handlers.CertificateTemplate)
	studentHomeworkRouter(api, handlers.StudentHomework)
	studentCertificateRouter(api, handlers.StudentCertificate)

	return router
}

func courseRouter(
	api *gin.RouterGroup,
	h *v1.Handlers,
) {
	course := api.Group("/courses")
	{
		course.GET("", h.Course.GetCourses)
		course.GET("/:course_id", h.Course.GetCourseById)
		course.POST("", h.Course.AddCourse)
		course.PUT(":course_id", h.Course.UpdateCourseById)
		course.DELETE("/:course_id", h.Course.DeleteCourseById)
	}
}

func personRouter(
	api *gin.RouterGroup,
	h *v1.Handlers,
) {
	person := api.Group("/persons")
	{
		person.GET("", h.Person.GetPersons)
		person.GET("/:person_id", h.Person.GetPersonById)
		person.DELETE("/:person_id", h.Person.DeletePersonById)
		person.POST("", h.Person.AddNewPerson)
		person.PUT("/:person_id", h.Person.UpdatePersonById)

	}
}

func studentRouter(
	api *gin.RouterGroup,
	h *v1.Handlers,
) {
	student := api.Group("/students")
	{
		student.GET("", h.Student.GetStudents)
		student.GET("/:student_id", h.Student.GetStudentByStudentId)
		student.DELETE("/:student_id", h.Student.DeleteStudentByStudentId)
		student.POST("", h.Student.AddStudent)
		student.PUT("/:student_id", h.Student.UpdateStudentByStudentId)

	}
}

func mentorRouter(
	api *gin.RouterGroup,
	h *v1.Handlers,
) {
	mentor := api.Group("/mentors")
	{
		mentor.GET("", h.Mentor.GetMentors)
		mentor.GET("/:mentor_id", h.Mentor.GetMentorByMentorId)
		mentor.POST("", h.Mentor.AddMentor)
		mentor.DELETE("/:mentor_id", h.Mentor.DeleteMentorByMentorId)
		mentor.PUT("/:mentor_id", h.Mentor.UpdateMentorByMentorId)

	}
}

func mentorNoteRouter(
	api *gin.RouterGroup,
	h *v1.Handlers,
) {
	mentorNote := api.Group("/mentors/notes")
	{
		mentorNote.GET("", h.MentorNote.GetMentorNotes)
		mentorNote.GET("/:mentor_id", h.MentorNote.GetMentorNotesByMentorId)
		// mentorNote.GET("/:mentor_note_id", h.MentorNote.GetMentorNoteByMentorNoteId)
		mentorNote.POST("", h.MentorNote.AddMentorNote)
		mentorNote.PUT("/:mentor_note_id", h.MentorNote.UpdateMentorNoteByMentorNoteId)
		mentorNote.DELETE("/:mentor_note_id", h.MentorNote.DeleteMentorNoteByMentorNoteId)
	}
}

func studentNoteRouter(
	api *gin.RouterGroup,
	h *v1.Handlers,
) {
	studentNote := api.Group("/students/notes")
	{
		//studentNote.GET("", h.StudentNote.GetStudentNotes)
		studentNote.GET("/:student_id", h.StudentNote.GetStudentNotesByStudentId)
		// studentNote.GET("/:student_note_id", h.StudentNote.GetStudentNoteByStudentNoteId)
		studentNote.POST("", h.StudentNote.AddStudentNote)
		studentNote.PUT("/:student_note_id", h.StudentNote.UpdateStudentNoteByStudentNoteId)
		studentNote.DELETE("/:student_note_id", h.StudentNote.DeleteStudentNote)
	}
}

func studentNoteTypeRouter(
	api *gin.RouterGroup,
	h *v1.Handlers,
) {
	studentNoteType := api.Group("/students/notes/types")
	{
		studentNoteType.GET("/types", h.StudentNoteType.GetStudentNoteTypes)
		studentNoteType.GET("/types/:type_id", h.StudentNoteType.GetStudentNoteTypeById)
		studentNoteType.POST("", h.StudentNoteType.AddStudentNoteType)
		studentNoteType.PUT("/:student_note_type_id", h.StudentNoteType.UpdateStudentNoteTypeById)
		studentNoteType.DELETE("/:student_note_type_id", h.StudentNoteType.DeleteStudentNoteTypeById)
	}

}

func groupContactRouter(
	api *gin.RouterGroup,
	h *v1.Handlers,
) {
	groupContact := api.Group("/groups/contacts")
	{
		groupContact.GET("", h.GroupContact.GetGroupContacts)
		groupContact.GET("/:group_contact_id", h.GroupContact.GetGroupContactsByGroupId)
		groupContact.POST("", h.GroupContact.AddGroupContact)
		groupContact.PUT("/:group_contact_id", h.GroupContact.UpdateGroupContact)
		groupContact.DELETE("/:group_contact_id", h.GroupContact.DeleteGroupContact)
	}
}

func studentGroupRouter(
	api *gin.RouterGroup,
	h *v1.Handlers,
) {
	studentGroup := api.Group("/groups/students")
	{
		studentGroup.GET("", h.StudentGroup.GetStudentGroups)
		studentGroup.GET("/:student_group_id", h.StudentGroup.GetStudentGroupById)
		studentGroup.POST("", h.StudentGroup.AddStudentGroup)
		studentGroup.PUT("/:student_group_id", h.StudentGroup.UpdateStudentGroupById)
		studentGroup.DELETE("/:student_group_id", h.StudentGroup.DeleteStudentGroupById)

	}
}

func courseStatusRouter(
	api *gin.RouterGroup,
	h *v1.Handlers,
) {
	courseStatus := api.Group("/courses/statuses")
	{
		courseStatus.GET("", h.CourseStatus.GetCourseStatuses)
		courseStatus.GET("/:course_status_id", h.CourseStatus.GetCourseStatusById)
		courseStatus.POST("", h.CourseStatus.AddCourseStatus)
		courseStatus.PUT("/:course_status_id", h.CourseStatus.UpdateCourseStatusById)
		courseStatus.DELETE("/:course_status_id", h.CourseStatus.DeleteCourseStatusById)
	}
}

func courseLectureRouter(
	api *gin.RouterGroup,
	h *v1.Handlers,
) {
	courseLecture := api.Group("/courses/lectures")
	{
		courseLecture.GET("", h.CourseLecture.GetCourseLectures)
		courseLecture.GET("/:course_lecture_id", h.CourseLecture.GetAllCourseLectures)
		courseLecture.POST("", h.CourseLecture.AddCourseLecture)
		courseLecture.PUT("/:course_lecture_id", h.CourseLecture.UpdateCourseLectureById)
		courseLecture.DELETE("/:course_lecture_id", h.CourseLecture.DeleteCourseLectureById)
	}
}

func certificateTemplateRouter(
	api *gin.RouterGroup,
	h *v1.CertificateTemplateHandlers,
) {
	certificateTemplate := api.Group("/certificate-templates")
	{
		certificateTemplate.GET("", h.GetCertificateTemplates)
		certificateTemplate.GET("/:certificate_template_id", h.GetCertificateTemplateById)
		certificateTemplate.POST("", h.AddCertificateTemplate)
		certificateTemplate.PUT("/:certificate_template_id", h.UpdateCertificateTemplate)
		certificateTemplate.DELETE("/:certificate_template_id", h.DeleteCertificateTemplate)
	}
}

func homeworkRouter(
	api *gin.RouterGroup,
	h *v1.HomeworkHandlers,
) {
	homework := api.Group("/homeworks")
	{
		homework.GET("", h.GetHomeworks)
		homework.GET("/:homework_id", h.GetHomeworkById)
		homework.POST("", h.AddHomework)
		homework.PUT("/:homework_id", h.UpdateHomework)
		homework.DELETE("/:homework_id", h.DeleteHomework)
	}

	api.GET("/lectures/:lecture_id/homeworks", h.GetHomeworksByLectureId)
}

func lectureRouter(
	api *gin.RouterGroup,
	h *v1.LectureHandlers,
) {
	lecture := api.Group("/lectures")
	{
		lecture.GET("", h.GetLectures)
		lecture.GET("/:lecture_id", h.GetLectureById)
		lecture.POST("", h.AddLecture)
		lecture.PUT("/:lecture_id", h.UpdateLecture)
		lecture.DELETE("/:lecture_id", h.DeleteLecture)
	}
}

func studentCertificateRouter(
	api *gin.RouterGroup,
	h *v1.StudentCertificateHandlers,
) {
	studentCertificate := api.Group("/student-certificates")
	{
		studentCertificate.GET("", h.GetStudentCertificates)
		studentCertificate.GET("/:student_certificate_id", h.GetStudentCertificateById)
		studentCertificate.POST("", h.AddStudentCertificate)
		studentCertificate.PUT("/:student_certificate_id", h.UpdateStudentCertificate)
		studentCertificate.DELETE("/:student_certificate_id", h.DeleteStudentCertificate)
	}

	api.GET("/students/:student_id/certificates", h.GetStudentCertificatesByStudentId)

	api.GET("/courses/:course_id/certificates", h.GetStudentCertificatesByCourseId)
}

func studentHomeworkRouter(
	api *gin.RouterGroup,
	h *v1.StudentHomeworkHandlers,
) {
	studentHomework := api.Group("/student-homeworks")
	{
		studentHomework.GET("", h.GetStudentHomeworks)
		studentHomework.GET("/:student_homework_id", h.GetStudentHomeworkById)
		studentHomework.POST("", h.AddStudentHomework)
		studentHomework.PUT("/:student_homework_id", h.UpdateStudentHomework)
		studentHomework.DELETE("/:student_homework_id", h.DeleteStudentHomework)
	}

	api.GET("/students/:student_id/homeworks", h.GetStudentHomeworksByStudentId)
}
