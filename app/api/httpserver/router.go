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
	studentRouter(api, handlers)
	mentorRouter(api, handlers)
	mentorNoteRouter(api, handlers)
	studentNoteRouter(api, handlers)
	studentNoteTypeRouter(api, handlers)
	groupContactRouter(api, handlers)
	studentGroupRouter(api, handlers)
	courseStatusRouter(api, handlers)
	courseLectureRouter(api, handlers)

	return router
}

func courseRouter(
	api *gin.RouterGroup,
	h *v1.Handlers,
) {
	course := api.Group("/course")
	{
		course.GET("", h.Course.GetCourses)
		course.GET("/:course_id", h.Course.GetCourseById)
		course.POST("", h.Course.AddCourse)
		course.PUT(":course_id", h.Course.EditCourseById)
		course.DELETE("/:course_id", h.Course.DeleteCourseByID)

	}
}

func personRouter(
	api *gin.RouterGroup,
	h *v1.Handlers,
) {
	person := api.Group("/person")
	{
		person.GET("", h.Person.GetPersons)
		person.GET("/:person_id", h.Person.GetPersonById)
		person.DELETE("/:person_id", h.Person.AddNewPerson)
		person.POST("", h.Person.AddNewPerson)
		person.PUT("/:person_id", h.Person.EditPersonById)

	}
}

func studentRouter(
	api *gin.RouterGroup,
	h *v1.Handlers,
) {
	student := api.Group("/student")
	{
		student.GET("", h.Student.GetStudents)
		student.GET("/:student_id", h.Student.GetStudentByStudentId)
		student.DELETE("/:student_id", h.Student.DeleteStudentByStudentId)
		student.POST("", h.Student.AddStudent)
		student.PUT("/:student_id", h.Student.EditStudentByStudentId)

	}
}

func mentorRouter(
	api *gin.RouterGroup,
	h *v1.Handlers,
) {
	mentor := api.Group("/mentor")
	{
		mentor.GET("", h.Mentor.GetMentors)
		mentor.GET("/:mentor_id", h.Mentor.GetMentorByMentorId)
		mentor.POST("", h.Mentor.AddMentor)
		mentor.DELETE("/:mentor_id", h.Mentor.RemoveMentorByMentorId)
		mentor.PUT("/:mentor_id", h.Mentor.EditMentorByMentorId)

	}
}

func mentorNoteRouter(
	api *gin.RouterGroup,
	h *v1.Handlers,
) {
	mentorNote := api.Group("/mentor/note")
	{
		mentorNote.GET("", h.MentorNote.GetMentorNotes)
		mentorNote.GET("/:mentor_id", h.MentorNote.GetMentorNotesByMentorId)
		mentorNote.GET("/:mentor_note_id", h.MentorNote.GetMentorNoteByMentorNoteId)
		mentorNote.POST("", h.MentorNote.AddNewMentorNote)
		mentorNote.PUT("/:mentor_note_id", h.MentorNote.EditMentorNoteByMentorNoteId)
		mentorNote.DELETE("/:mentor_note_id", h.MentorNote.DeleteMentorNoteByMentorNoteId)
	}
}

func studentNoteRouter(
	api *gin.RouterGroup,
	h *v1.Handlers,
) {
	studentNote := api.Group("/student/note")
	{
		studentNote.GET("", h.StudentNote.GetStudentNotes)
		studentNote.GET("/:student_id", h.StudentNote.GetStudentNotesByStudentId)
		studentNote.GET("/:student_note_id", h.StudentNote.GetStudentNoteByStudentNoteId)
		studentNote.POST("", h.StudentNote.AddStudentNote)
		studentNote.PUT("/:student_note_id", h.StudentNote.EditStudentNoteByStudentNoteId)
		studentNote.DELETE("/:student_note_id", h.StudentNote.DeleteStudentNote)
	}
}

func studentNoteTypeRouter(
	api *gin.RouterGroup,
	h *v1.Handlers,
) {
	studentNoteType := api.Group("/student/note_type")
	{
		studentNoteType.GET("", h.StudentNoteType.GetStudentNoteTypes)
		studentNoteType.GET("/:student_note_type_id", h.StudentNoteType.GetStudentNoteTypeById)
		studentNoteType.POST("", h.StudentNoteType.AddStudentNoteType)
		studentNoteType.PUT("/:student_note_type_id", h.StudentNoteType.EditStudentNoteTypeById)
		studentNoteType.DELETE("/:student_note_type_id", h.StudentNoteType.DeleteStudentNoteTypeById)
	}

}

func groupContactRouter(
	api *gin.RouterGroup,
	h *v1.Handlers,
) {
	groupContact := api.Group("/group/contact")
	{
		groupContact.GET("", h.GroupContact.GetGroupContacts)
		groupContact.GET("/:group_contact_id", h.GroupContact.GetGroupContactById)
		groupContact.POST("", h.GroupContact.AddGroupContact)
		groupContact.PUT("/:group_contact_id", h.GroupContact.UpdateGroupContact)
		groupContact.DELETE("/:group_contact_id", h.GroupContact.DeleteGroupContact)
	}
}

func studentGroupRouter(
	api *gin.RouterGroup,
	h *v1.Handlers,
) {
	studentGroup := api.Group("/group/students")
	{
		studentGroup.GET("", h.StudentGroup.GetStudentGroups)
		studentGroup.GET("/:student_group_id", h.StudentGroup.GetStudentGroupById)
		studentGroup.POST("", h.StudentGroup.AddStudentGroup)
		studentGroup.PUT("/:student_group_id", h.StudentGroup.AddStudentGroup)
		studentGroup.DELETE("/:student_group_id", h.StudentGroup.DeleteStudentGroup)

	}
}

func courseStatusRouter(
	api *gin.RouterGroup,
	h *v1.Handlers,
) {
	courseStatus := api.Group("/course/status")
	{
		courseStatus.GET("", h.CourseStatus.GetCourseStatuses)
		courseStatus.GET("/:course_status_id", h.CourseStatus.GetCourseStatusById)
		courseStatus.POST("", h.CourseStatus.AddCourseStatus)
		courseStatus.PUT("/:course_status_id", h.CourseStatus.EditCourseStatusById)
		courseStatus.DELETE("/:course_status_id", h.CourseStatus.DeleteCourseStatusById)
	}
}

func courseLectureRouter(
	api *gin.RouterGroup,
	h *v1.Handlers,
) {
	courseLecture := api.Group("/course/lecture")
	{
		courseLecture.GET("", h.CourseLecture.GetCourseLectures)
		courseLecture.GET("/course_lecture_id", h.CourseLecture.GetCourseLectureById)
		courseLecture.POST("", h.CourseLecture.AddCourseLecture)
		courseLecture.PUT("/:course_lecture_id", h.CourseLecture.EditCourseLectureById)
		courseLecture.DELETE("/:course_lecture_id", h.CourseLecture.DeleteCourseLectureById)
	}
}
