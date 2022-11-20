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

// GetHomeworks
// @Summary get all items in the homework list
// @Description возвращает все домашние работы
// @ID get-all-homeworks
// @Tags homeworks
// @Produce json
// @Success 200 {object} model.HomeworkList
// @Failure 500 {object} model.ResponseMessage
// @Router /homeworks [get]
func (h *HomeworkHandlers) GetHomeworks(c *gin.Context) {
}

// GetHomeworksByLectureId
// @Summary get items in the homework list by lecture ID
// @Description возвращает домашнюю работы по указанному id лекции
// @ID get-homeworks-by-lecture-id
// @Tags homeworks
// @Produce json
// @Param lecture_id path string true "lecture id"
// @Success 200 {object} model.HomeworkList
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /lectures/{lecture_id}/homeworks [get]
func (h *HomeworkHandlers) GetHomeworksByLectureId(c *gin.Context) {
}

// GetHomeworkById
// @Summary get a homework by ID
// @Description возвращает домашнюю работу с указанным id
// @ID get-homework-by-id
// @Tags homeworks
// @Produce json
// @Param id path string true "homework id"
// @Success 200 {object} model.Homework
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /homeworks/{id} [get]
func (h *HomeworkHandlers) GetHomeworkById(c *gin.Context) {
}

// AddHomework
// @Summary add new homework to the homework list
// @Description добавляет новую домашнюю работу
// @ID create-homework
// @Tags homeworks
// @Produce json
// @Consume json
// @Param homework body model.Homework true "homework"
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /homeworks [post]
func (h *HomeworkHandlers) AddHomework(c *gin.Context) {
}

// UpdateHomework
// @Summary update a homework by ID
// @Description изменяет домашнюю работу с указанным id
// @ID update-homework-by-id
// @Tags homeworks
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

// DeleteHomework
// @Summary delete a homework by ID
// @Description удаляет домашнюю работу с указанным id
// @ID delete-homework-by-id
// @Tags homeworks
// @Produce json
// @Param id path string true "homework id"
// @Success 200 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /homeworks/{id} [delete]
func (h *HomeworkHandlers) DeleteHomework(c *gin.Context) {
}
