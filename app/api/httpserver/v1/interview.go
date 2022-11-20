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

// @Summary add new item to the interview list
// @ID create-interview
// @Tags interviews
// @Produce json
// @Consume json
// @Param interview body model.Interview true "interview"
// @Success 201 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /interviews [post]
func AddInterview(c *gin.Context) {}

// @Summary get all items in the interview list
// @ID get-all-interviews
// @Tags interviews
// @Produce json
// @Success 200 {object} []model.Interview
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /interviews [get]
func GetInterviews(c *gin.Context) {}

// @Summary get a interview by ID
// @ID get-interview-by-id
// @Tags interviews
// @Produce json
// @Param id path string true "interview id"
// @Success 200 {object} model.Interview
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /interview/{id} [get]
func GetInterviewById(c *gin.Context) {}

// @Summary update a interview by ID
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
// @Router /interview/{id} [post]
func UpdateInterviewById(c *gin.Context) {}

// @Summary put a interview by ID
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

// @Summary delete a interview by ID
// @ID delete-interview-by-id
// @Tags interviews
// @Param id path string true "interview id"
// @Produce json
// @Success 200 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /interview/{id} [delete]
func DeleteInterviewById(c *gin.Context) {}
