package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type InterviewHandlers struct {
	lg      logger.Logger
	usecase usecase.InterviewUsecaseInterface
}

func NewInterview(lg logger.Logger, u usecase.InterviewUsecaseInterface) *InterviewHandlers {
	return &InterviewHandlers{
		lg:      lg,
		usecase: u,
	}
}

// AddInterview
// @Summary add new item to the interview list
// @Description добавляет новое интервью новое интервью
// @ID create-interview
// @Tags interviews
// @Produce json
// @Consume json
// @Param interview body model.Interview true "interview"
// @Success 201 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /interview [post]
func AddInterview(c *gin.Context) {}

// GetInterviews
// @Summary get all items in the interview list
// @Description возвращает все интервью
// @ID get-all-interviews
// @Tags interviews
// @Produce json
// @Success 200 {object} []model.Interview
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /interview [get]
func GetInterviews(c *gin.Context) {}

// GetInterviewById
// @Summary get a interview by ID
// @Description возвращает интервью с указанным id
// @ID get-interview-by-id
// @Tags interviews
// @Produce json
// @Param id path string true "interview id"
// @Success 200 {object} model.Interview
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /interview/{id} [get]
func GetInterviewById(c *gin.Context) {}

// PutInterviewById
// @Summary put a interview by ID
// @Description изменяет интервью с указанным id
// @ID update-interview-by-id
// @Tags interviews
// @Param id path string true "interview id"
// @Param interview body model.Interview true "interview"
// @Produce json
// @Consume json
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /interview/{id} [put]
func PutInterviewById(c *gin.Context) {}

// DeleteInterviewById
// @Summary delete a interview by ID
// @Description удаляет интервью с указанным id
// @ID delete-interview-by-id
// @Tags interviews
// @Param id path string true "interview id"
// @Produce json
// @Success 200 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /interview/{id} [delete]
func DeleteInterviewById(c *gin.Context) {}
