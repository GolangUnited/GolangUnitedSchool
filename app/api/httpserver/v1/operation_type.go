package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type OperationTypeHandlers struct {
	lg logger.Logger
}

func NewOperationType(lg logger.Logger) *OperationTypeHandlers {
	return &OperationTypeHandlers{lg: lg}
}

// AddOperationType
// @Summary add new item to the operation types list
// @Description добавляет новый тип операции
// @ID create-operation-type
// @Tags operationTypes
// @Produce json
// @Consume json
// @Param operation-type body model.OperationType true "operation-type"
// @Success 201 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /operation-type [post]
func (h *OperationTypeHandlers) AddOperationType(c *gin.Context) {}

// GetOperationTypes
// @Summary get all items in the operation type list
// @Description возвращает все типы операций
// @ID get-all-operation-types
// @Tags operationTypes
// @Produce json
// @Success 200 {object} []model.OperationType
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /operation-type [get]
func (h *OperationTypeHandlers) GetOperationTypes(c *gin.Context) {}

// GetOperationTypeById
// @Summary get a operation type by ID
// @Description возвращает тип операции с указанным id
// @ID get-operation-type-by-id
// @Tags operationTypes
// @Produce json
// @Param id path string true "operation-type id"
// @Success 200 {object} model.OperationType
// @Failure 404 {object} model.ResponseMessage
// @Router /operation-type/{id} [get]
func (h *OperationTypeHandlers) GetOperationTypeById(c *gin.Context) {}

// UpdateOperationTypeById
// @Summary put a operation type by ID
// @Description изменяет тип указанному операции с указанным id
// @ID update-operation-type-by-id
// @Tags operationTypes
// @Param id path string true "operation type id"
// @Param operation_type body model.OperationType true "role"
// @Produce json
// @Consume json
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /operation-type/{id} [put]
func (h *OperationTypeHandlers) UpdateOperationTypeById(c *gin.Context) {}

// DeleteOperationTypeById
// @Summary delete a operation type by ID
// @Description  удаляет тип операции с указанным id
// @ID delete-operation-type-by-id
// @Tags operationTypes
// @Param id path string true "operation type id"
// @Produce json
// @Success 200 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /operation-type/{id} [delete]
func (h *OperationTypeHandlers) DeleteOperationTypeById(c *gin.Context) {}
