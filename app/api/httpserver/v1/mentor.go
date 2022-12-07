package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type MentorHandlers struct {
	lg      logger.Logger
	useCase usecase.MentorUseCaseInterface
}

func NewMentorHandler(
	lg logger.Logger,
	useCase usecase.MentorUseCaseInterface,
) *MentorHandlers {
	return &MentorHandlers{
		lg:      lg,
		useCase: useCase,
	}
}

// GetMentors
// @Summary get all mentors
// @Description получить список всех менторов
// @ID get-all-mentors
// @Tags mentors
// @Produce json
// @Success 200 {object} []model.Mentor
// @Failure 500 {object} model.ResponseMessage
// @Router /mentors [get]
func (h *MentorHandlers) GetMentors(c *gin.Context) {}

// GetMentorByMentorId
// @Summary get mentor by mentor id
// @Description получить данные о менторе по его id, выдает структуру person
// @ID get-mentor-by-id
// @Tags mentors
// @Produce json
// @Success 200 {object} model.Person
// @Failure 500 {object} model.ResponseMessage
// @Router /mentors/{mentor_id} [get]
func (h *MentorHandlers) GetMentorByMentorId(c *gin.Context) {}

// AddMentor
// @Summary add mentor
// @Description добавить пользователя в группу менторов по его person_id
// @ID add-mentor
// @Tags mentors
// @Produce json
// @Param mentor body model.UpdateMentor true "mentor"
// @Success 200 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /mentors/{person_id} [post]
func (h *MentorHandlers) AddMentor(c *gin.Context) {}

// PutMentorByMentorId
// @Summary update mentor by mentor id
// @Description изменить пользователя группы "менторы"
// @ID put-mentor
// @Tags mentors
// @Produce json
// @Param id path string true "mentor_id"
// @Param mentor body model.UpdateMentor true "mentor"
// @Success 200 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /mentors/{person_id} [put]
func (h *MentorHandlers) PutMentorByMentorId(c *gin.Context) {}

// DeleteMentorByMentorId
// @Summary delete mentor by mentor id
// @Description удалить пользователя из группы менторов //
// @ID delete-mentor-by-id
// @Tags mentors
// @Produce json
// @Param id path string true "mentor_id"
// @Success 200 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /mentors/{mentor_id} [delete]
func (h *MentorHandlers) DeleteMentorByMentorId(c *gin.Context) {}
