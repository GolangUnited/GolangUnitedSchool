package httpserver

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lozovoya/GolangUnitedSchool/app/api/httpserver/v1"
)

func NewRouter(
	handlers *v1.Handlers,
) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")
	courseRouter(api, handlers)
	personRouter(api, handlers)
	lectureRouter(api, handlers.Lecture)
	homeworkRouter(api, handlers.Homework)
	projectRouter(api, handlers.Project)
	certificateTemplateRouter(api, handlers.CertificateTemplate)
	studentHomeworkRouter(api, handlers.StudentHomework)
	studentCertificateRouter(api, handlers.StudentCertificate)

	return router
}

func courseRouter(
	api *gin.RouterGroup,
	h *v1.Handlers,
) {
	course := api.Group("/course")
	{
		course.GET("")
		course.GET("/:course_id")
	}
}

func personRouter(
	api *gin.RouterGroup,
	h *v1.Handlers,
) {
	person := api.Group("/person")
	{
		person.GET("")
		person.GET("/:person_id")
		person.DELETE("/")
		person.DELETE("/:person_id")
	}
}

func certificateTemplateRouter(
	api *gin.RouterGroup,
	h *v1.CertificateTemplateHandlers,
) {
	certificateTemplate := api.Group("/certificate_template")
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
	homework := api.Group("/homework")
	{
		homework.GET("", h.GetHomeworks)
		homework.GET("/:homework_id", h.GetHomeworkById)
		homework.POST("", h.AddHomework)
		homework.PUT("/:homework_id", h.UpdateHomework)
		homework.DELETE("/:homework_id", h.DeleteHomework)
	}

	api.GET("/lecture/:lecture_id/homework", h.GetHomeworksByLectureId)
}

func lectureRouter(
	api *gin.RouterGroup,
	h *v1.LectureHandlers,
) {
	lecture := api.Group("/lecture")
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
	studentCertificate := api.Group("/student_certificate")
	{
		studentCertificate.GET("", h.GetStudentCertificates)
		studentCertificate.GET("/:student_certificate_id", h.GetStudentCertificateById)
		studentCertificate.POST("", h.AddStudentCertificate)
		studentCertificate.PUT("/:student_certificate_id", h.UpdateStudentCertificate)
		studentCertificate.DELETE("/:student_certificate_id", h.DeleteStudentCertificate)
	}

	api.GET("/student/:student_id/certificate", h.GetStudentCertificatesByStudentId)

	api.GET("/course/:course_id/certificate", h.GetStudentCertificatesByCourseId)
}

func studentHomeworkRouter(
	api *gin.RouterGroup,
	h *v1.StudentHomeworkHandlers,
) {
	studentHomework := api.Group("/student_homework")
	{
		studentHomework.GET("", h.GetStudentHomeworks)
		studentHomework.GET("/:student_homework_id", h.GetStudentHomeworkById)
		studentHomework.POST("", h.AddStudentHomework)
		studentHomework.PUT("/:student_homework_id", h.UpdateStudentHomework)
		studentHomework.DELETE("/:student_homework_id", h.DeleteStudentHomework)
	}

	api.GET("/student/:student_id/homework", h.GetStudentHomeworksByStudentId)

}

func projectRouter(api *gin.RouterGroup, h *v1.ProjectHandlers) {
	project := api.Group("/project")
	{
		project.GET("", h.GetProjects)
		project.GET("/:project_id", h.GetProjectByID)
		project.POST("", h.AddProject)
		project.PUT("/:project_id", h.UpdateProject)
		project.DELETE("/:project_id", h.DeleteProject)
	}
}
