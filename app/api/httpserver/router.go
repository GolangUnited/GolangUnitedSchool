package httpserver

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lozovoya/GolangUnitedSchool/app/api/httpserver/v1"
)

func NewRouter(
	courseHandler *v1.CourseHandlers,
	personHandler *v1.PersonHandlers,
	studentHandler *v1.StudentHandlers,
	mentorHandler *v1.MentorHandlers,
	mentorNoteHandler *v1.MentorNoteHandlers,
	studentNoteHandler *v1.StudentNoteHandlers,
) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")
	courseRouter(api, courseHandler)
	personRouter(api, personHandler)
	studentRouter(api, studentHandler)
	mentorRouter(api, mentorHandler)
	mentorNoteRouter(api, mentorNoteHandler)
	studentNoteRouter(api, studentNoteHandler)

	return router
}

func courseRouter(
	api *gin.RouterGroup,
	h *v1.CourseHandlers,
) {
	course := api.Group("/course")
	{
		course.GET("", h.SearchCourse)
		course.GET("/:course_id", h.GetCourseByID)
		course.POST("", h.CreateCourse)
		course.PATCH("", h.EditCourse)
		course.PATCH("/:course_id", h.EditCourseByID)
		course.PUT("", h.AddCourse)
		course.DELETE("", h.DeleteCourse)
		course.DELETE("/:course_id", h.DeleteCourseByID)
	}
}

func personRouter(
	api *gin.RouterGroup,
	h *v1.PersonHandlers,
) {
	person := api.Group("/person")
	{
		person.GET("", h.SearchPerson)
		person.GET("/:person_id", h.GetPersonByID)
		person.DELETE("/", h.DeletePerson)
		person.DELETE("/:person_id", h.DeletePersonByID)
		person.POST("/", h.AddNewPerson)
		person.PUT("/:person_id", h.EditPersonById)
		person.PUT("/", h.EditPerson)
	}
}

func studentRouter(
	api *gin.RouterGroup,
	h *v1.StudentHandlers,
) {
	student := api.Group("/student")
	{
		student.GET("", h.SearchStudent)
		student.GET(":person_id", h.GetStudentById)
		student.DELETE("", h.DeleteStudent)
		student.DELETE(":student_id", h.DeleteStudentById)
		student.POST("", h.AddStudent)
		student.POST(":person_id", h.AddStudentByPersonId)

	}
}

func mentorRouter(
	api *gin.RouterGroup,
	h *v1.MentorHandlers,
) {
	mentor := api.Group("/mentor")
	{
		mentor.GET("", h.GetAllMentors)
		mentor.GET(":mentor_id", h.GetMentorById)
		mentor.GET("", h.SearchMentor)
		mentor.POST("", h.AddMentorByName)
		mentor.DELETE("mentor_id", h.RemoveMentorById)
		mentor.DELETE("", h.RemoveMentorByName)
	}
}

func mentorNoteRouter(
	api *gin.RouterGroup,
	h *v1.MentorNoteHandlers,
) {
	mentorNote := api.Group("/mentor/note")
	{
		mentorNote.GET("", h.GetMentorNotes)
		mentorNote.GET(":mentor_note_id", h.GetMentorNoteById)
		mentorNote.POST("", h.AddNewMentorNote)
		mentorNote.PUT(":mentor_note_id", h.UpdateMentorNotebyId)
		mentorNote.DELETE(":mentor_note_id", h.DeleteMentorNoteById)
	}
}

func studentNoteRouter(
	api *gin.RouterGroup,
	h *v1.StudentNoteHandlers,
) {
	studentNote := api.Group("/student/note")
	{
		studentNote.GET(":student_id", h.GetStudentNotesByStudentId)
		studentNote.GET(":student_note_id", h.GetStudentNoteById)
		studentNote.POST("", h.AddStudentNote)
		studentNote.PUT("", h.UpdateStudentNote)
		studentNote.DELETE("student_note_id", h.DeleteStudentNote)
	}
}
