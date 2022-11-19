package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type HomeworkHandlers struct {
	lg      logger.Logger
	usecase usecase.HomeworkUsecaseInterface
}

func NewHomeworkHandler(
	lg logger.Logger,
	usecase usecase.HomeworkUsecaseInterface,
) *HomeworkHandlers {
	return &HomeworkHandlers{
		lg:      lg,
		usecase: usecase,
	}
}

// @Summary get all items in the homework list
// @ID get-all-homeworks
// @Tags courses
// @Produce json
// @Success 200 {object} model.HomeworkList
// @Failure 500 {object} model.ResponseMessage
// @Router /homeworks [get]
func (h *HomeworkHandlers) GetHomeworks(c *gin.Context) {
}

// @Summary get items in the homework list by lecture ID
// @ID get-homeworks-by-lecture-id
// @Tags courses
// @Produce json
// @Param lecture_id path string true "lecture id"
// @Success 200 {object} model.HomeworkList
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /lectures/{lecture_id}/homeworks [get]
func (h *HomeworkHandlers) GetHomeworksByLectureId(c *gin.Context) {
}

// @Summary get a homework by ID
// @ID get-homework-by-id
// @Tags courses
// @Produce json
// @Param id path string true "homework id"
// @Success 200 {object} model.Homework
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /homeworks/{id} [get]
func (h *HomeworkHandlers) GetHomeworkById(c *gin.Context) {
}

// @Summary add new homework to the homework list
// @ID create-homework
// @Tags courses
// @Produce json
// @Consume json
// @Param homework body model.Homework true "homework"
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /homeworks [post]
func (h *HomeworkHandlers) AddHomework(c *gin.Context) {
}

// @Summary update a homework by ID
// @ID update-homework-by-id
// @Tags courses
// @Produce json
// @Consume json
// @Param id path string true "homework id"
// @Param homework body model.Homework true "homework"
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /homeworks/{id} [put]
func (h *HomeworkHandlers) UpdateHomework(c *gin.Context) {
}

// @Summary delete a homework by ID
// @ID delete-homework-by-id
// @Tags courses
// @Produce json
// @Param id path string true "homework id"
// @Success 200 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /homeworks/{id} [delete]
func (h *HomeworkHandlers) DeleteHomework(c *gin.Context) {
}
