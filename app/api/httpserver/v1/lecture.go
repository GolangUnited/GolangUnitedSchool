package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type LectureHandlers struct {
	lg      logger.Logger
	usecase usecase.LectureUsecaseInterface
}

func NewLectureHandler(
	lg logger.Logger,
	usecase usecase.LectureUsecaseInterface,
) *LectureHandlers {
	return &LectureHandlers{
		lg:      lg,
		usecase: usecase,
	}
}

// @Summary get all items in the lecture list
// @ID get-all-lectures
// @Tags courses
// @Produce json
// @Success 200 {object} model.LectureList
// @Failure 500 {object} model.ResponseMessage
// @Router /lectures [get]
func (h *LectureHandlers) GetLectures(c *gin.Context) {
}

// @Summary get a lecture by ID
// @ID get-lecture-by-id
// @Tags courses
// @Produce json
// @Param id path string true "lecture id"
// @Success 200 {object} model.Lecture
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /lectures/{id} [get]
func (h *LectureHandlers) GetLectureById(c *gin.Context) {
}

// @Summary add new lecture to the lecture list
// @ID create-lecture
// @Tags courses
// @Produce json
// @Consume json
// @Param lecture body model.Lecture true "lecture"
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /lectures [post]
func (h *LectureHandlers) AddLecture(c *gin.Context) {
}

// @Summary update a lecture by ID
// @ID update-lecture-by-id
// @Tags courses
// @Produce json
// @Consume json
// @Param id path string true "lecture id"
// @Param lecture body model.Lecture true "lecture"
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /lectures/{id} [put]
func (h *LectureHandlers) UpdateLecture(c *gin.Context) {
}

// @Summary delete a lecture by ID
// @ID delete-lecture-by-id
// @Tags courses
// @Produce json
// @Param id path string true "lecture id"
// @Success 200 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /lectures/{id} [delete]
func (h *LectureHandlers) DeleteLecture(c *gin.Context) {
}
