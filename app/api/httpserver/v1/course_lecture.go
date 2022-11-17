package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type CourseLectureHandlers struct {
	lg      logger.Logger
	useCase usecase.CourseStatusUseCaseInterface
}

func NewCourseLectureHandler(
	lg logger.Logger,
	useCase usecase.CourseLectureUseCaseInterface,
) *CourseLectureHandlers {
	return &CourseLectureHandlers{
		lg:      lg,
		useCase: useCase,
	}
}

// GetCourseLectures
// @ Summary get all course lectures
// @ Description возвращает список всех лекций всех курсов
// @ ID get-all-course-lectures
// @ Tags courses, courseLectures
// @ Produce json
// @ Success 200 {object} model.CourseLecturesListDto
// @ Failure 500 {object} model.ResponseMessage
// @ Router /courses/lectures [get]
func (h *CourseLectureHandlers) GetCourseLectures(c *gin.Context) {}

// GetAllCourseLectures
// @Summary get a course lecture by id
// @ID get-course-lectures-by-course-id
// @ Tags courses, courseLectures
// @ Description возвращает все лекции определенного курса
// @Produce json
// @Param id path string true "course_id"
// @Success 200 {object} model.CourseLecturesListDto
// @Failure 404 {object} model.ResponseMessage
// @Router /courses/{course_id}/lectures [get]
func (h *CourseLectureHandlers) GetAllCourseLectures(c *gin.Context) {}

// GetCourseLecture
// @Summary get concrete lecture from course
// @Description возвращает лекцию с указанным lecture_id с курса course_id
// @ID get--lecture-by-id-from-course
// @ Tags courses, courseLectures
// @Produce json
// @Param course_id path string true "course_id"
// @Param lecture_id path string true "lecture_id"
// @Success 200 {object} model.CourseLecture
// @Failure 404 {object} model.ResponseMessage
// @Router /courses/{course_id}/lectures/{lecture_id} [get]
func (h *CourseLectureHandlers) GetCourseLecture(c *gin.Context) {}

// AddCourseLecture
// @Summary add new course lecture
// @ID add-course-lecture-by-id
// @Description добавить новую лекцию на курс
// @Tags courses, courseLectures
// @Produce json
// @Consume json
// @Param course body model.CourseLecture true "course_lecture"
// @Success 201 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /courses/lectures [post]
func (h *CourseLectureHandlers) AddCourseLecture(c *gin.Context) {}

// UpdateCourseLectureById
// @Summary update course lecture by id
// @Description отредактировать лекцию курса с указанным lecture_id
// @ID update-course-lecture-by-id
// @Tags courses, courseLectures
// @Param id path string true "lecture_id"
// @Param course lecture body model.CourseLecture true "course_lecture"
// @Produce json
// @Consume json
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Router /courses/lectures/{lecture_id} [put]
func (h *CourseLectureHandlers) UpdateCourseLectureById(c *gin.Context) {}

// DeleteCourseLectureById
// @Summary delete a course lecture by id
// @Description удалить лекцию с указанным lecture_id
// @ID delete-course-lecture-by-id
// @Tags courses, courseLectures
// @Param id path string true "course_lecture_id"
// @Produce json
// @Success 200 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /courses/lectures/{lecture_id} [delete]
func (h *CourseLectureHandlers) DeleteCourseLectureById(c *gin.Context) {}
