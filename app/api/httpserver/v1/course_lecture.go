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
// @Summary get all course lectures
// @ID get-all-course-lectures
// @Tags course_lecture
// @Produce json
// @Success 200 {object} model.CourseLectureListDto
// @Failure 500 {object} model.ResponseMessage
// @Router /course/lecture [get]
func (h *CourseLectureHandlers) GetCourseLectures(c *gin.Context) {}

// GetCourseLectureById
// @Summary get a course lecture by id
// @ID get-course-lecture-by-id
// @Tags course_lecture
// @Produce json
// @Param id path string true "course_lecture_id"
// @Success 200 {object} model.CourseLecture
// @Failure 404 {object} model.ResponseMessage
// @Router /course/lecture/{course_lecture_id} [get]
func (h *CourseLectureHandlers) GetCourseLectureById(c *gin.Context) {}

// AddCourseLecture
// @ID add-course-lecture-by-id
// @Tags course_lecture
// @Produce json
// @Consume json
// @Param course body model.CourseLecture true "course lecture"
// @Success 201 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /course/lecture [post]
func (h *CourseLectureHandlers) AddCourseLecture(c *gin.Context) {}

// UpdateCourseLectureById
// @Summary update course lecture by id
// @ID update-course-lecture-by-id
// @Tags course_lecture
// @Param id path string true "course_lecture_id"
// @Produce json
// @Consume json
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Router /course/lecture/{course_lecture_id} [put]
func (h *CourseLectureHandlers) UpdateCourseLectureById(c *gin.Context) {}

// DeleteCourseLectureById
// @Summary delete a course lecture by id
// @ID delete-course-lecture-by-id
// @Tags course_lecture
// @Param id path string true "course_lecture_id"
// @Produce json
// @Success 200 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /course/lecture/{course_lecture_id} [delete]
func (h *CourseLectureHandlers) DeleteCourseLectureById(c *gin.Context) {}
