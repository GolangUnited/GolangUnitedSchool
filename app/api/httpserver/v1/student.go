package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type StudentHandlers struct {
	lg      logger.Logger
	usecase usecase.StudentUsecaseInterface
}

func NewStudentHandler(
	lg logger.Logger,
	usecase usecase.StudentUsecaseInterface,
) *StudentHandlers {
	return &StudentHandlers{
		lg:      lg,
		usecase: usecase,
	}
}

func (h *StudentHandlers) AddStudent(c *gin.Context) {
}
func (h *StudentHandlers) AddStudentByPersonId(c *gin.Context) {
}
func (h *StudentHandlers) DeleteStudent(c *gin.Context) {
}
func (h *StudentHandlers) DeleteStudentById(c *gin.Context) {
}
func (h *StudentHandlers) SearchStudent(c *gin.Context) {
}
func (h *StudentHandlers) GetStudentById(c *gin.Context) {

}
