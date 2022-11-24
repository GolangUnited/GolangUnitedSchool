package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type ProjectHandlers struct {
	lg      logger.Logger
	usecase usecase.ProjectUsecaseInterface
}

func NewProjectHandler(
	lg logger.Logger,
	usecase usecase.ProjectUsecaseInterface,
) *ProjectHandlers {
	return &ProjectHandlers{
		lg:      lg,
		usecase: usecase,
	}
}

// GetProjects
// @Summary get all items in the project list
// @Description возвращает все проекты
// @ID get-all-projects
// @Tags projects
// @Produce json
// @Success 200 {object} model.ProjectList
// @Failure 500 {object} model.ResponseMessage
// @Router /projects [get]
func (h *ProjectHandlers) GetProjects(c *gin.Context) {
}

// GetProjectsByCourseId
// @Summary get items in the project list by course ID
// @Description возвращает проекты по указанному id курса
// @ID get-projects-by-course-id
// @Tags projects
// @Produce json
// @Param course_id path string true "course id"
// @Success 200 {object} model.ProjectList
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /courses/{course_id}/projects [get]
func (h *ProjectHandlers) GetProjectsByCourseId(c *gin.Context) {
}

// GetProjectsByGroupId
// @Summary get items in the project list by group ID
// @Description возвращает проекты по указанному id группы
// @ID get-projects-by-group-id
// @Tags projects
// @Produce json
// @Param group_id path string true "group id"
// @Success 200 {object} model.ProjectList
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /groups/{group_id}/projects [get]
func (h *ProjectHandlers) GetProjectsByGroupId(c *gin.Context) {
}

// GetProjectById
// @Summary get a project by ID
// @Description возвращает проект с указанным id
// @ID get-project-by-id
// @Tags projects
// @Produce json
// @Param id path string true "project id"
// @Success 200 {object} model.Project
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /projects/{id} [get]
func (h *ProjectHandlers) GetProjectById(c *gin.Context) {
}

// CreateProject
// @Summary add new project to the project list
// @Description добавляет проект
// @ID create-project
// @Tags projects
// @Produce json
// @Consume json
// @Param project body model.Project true "project"
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /projects [post]
func (h *ProjectHandlers) CreateProject(c *gin.Context) {
}

// UpdateProject
// @Summary update a project by ID
// @Description изменяет проект с указанным id
// @ID update-project-by-id
// @Tags projects
// @Produce json
// @Consume json
// @Param id path string true "project id"
// @Param project body model.Project true "project"
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /projects/{id} [put]
func (h *ProjectHandlers) UpdateProject(c *gin.Context) {
}

// DeleteProject
// @Summary delete a project by ID
// @Description удалаяет проект с указанным id
// @ID delete-project-by-id
// @Tags projects
// @Produce json
// @Param id path string true "project id"
// @Success 200 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /projects/{id} [delete]
func (h *ProjectHandlers) DeleteProject(c *gin.Context) {
}
