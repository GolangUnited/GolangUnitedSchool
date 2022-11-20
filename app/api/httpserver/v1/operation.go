package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type OperationHandlers struct {
	lg logger.Logger
}

func NewOperation(lg logger.Logger) *OperationHandlers {
	return &OperationHandlers{lg: lg}
}

// AddOperation
// @Summary add new item to the operations list
// @Description добавляет новую операцию
// @ID create-operation
// @Tags operations
// @Produce json
// @Consume json
// @Param operation body model.Operation true "operation"
// @Success 201 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /log/operations [post]
func (h *OperationHandlers) AddOperatioLog(c *gin.Context) {}

// GetOperations
// @Summary get all items in the operations list
// @Description возвращает все операции
// @ID get-all-operations
// @Tags operations
// @Produce json
// @Success 200 {object} []model.Operation
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /log/operations [get]
func (h *OperationHandlers) GetOperations(c *gin.Context) {}

// GetOperationById
// @Summary get a operation by ID
// @Description возвращает операцию с указанным id
// @ID get-operation-by-id
// @Tags operations
// @Produce json
// @Param id path string true "operation id"
// @Success 200 {object} model.Operation
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /logs/operations/{id} [get]
func (h *OperationHandlers) GetOperationById(c *gin.Context) {}

// @Summary update a operation by ID
// @ID update-operation-by-id
// @Tags operations
// @Param id path string true "operation id"
// @Param operation body model.Operation true "role"
// @Produce json
// @Consume json
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /logs/operations/{id} [post]
func (h *OperationHandlers) UpdateOperationById(c *gin.Context) {}

// @Summary put a operation by ID
// @Description изменяет операцию по указанному id
// @ID update-operation-by-id
// @Tags operations
// @Param id path string true "operation id"
// @Param operation body model.Operation true "role"
// @Produce json
// @Consume json
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /logs/operations/{id} [put]
func (h *OperationHandlers) PutOperationById(c *gin.Context) {}

// DeleteOperationById
// @Summary delete a operation by ID
// @Description удаляет операцию по указанному id
// @ID delete-operation-by-id
// @Tags operations
// @Param id path string true "operation id"
// @Produce json
// @Success 200 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /logs/operations/{id} [delete]
func (h *OperationHandlers) DeleteOperationById(c *gin.Context) {}
