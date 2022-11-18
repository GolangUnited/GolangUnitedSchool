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

// @Summary add new item to the operation types list
// @ID create-operation-type
// @Tags operation-types
// @Produce json
// @Consume json
// @Param operation-type body model.OperationType true "operation-type"
// @Success 201 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /operation-type [post]
func (h *OperationTypeHandlers) AddOperatioLog(c *gin.Context) {}

// @Summary get all items in the operation type list
// @ID get-all-operation-types
// @Tags operation-types
// @Produce json
// @Success 200 {object} []model.OperationType
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /operation-type [get]
func (h *OperationTypeHandlers) GetOperationTypes(c *gin.Context) {}

// @Summary get a operation type by ID
// @ID get-operation-type-by-id
// @Tags operation-types
// @Produce json
// @Param id path string true "operation-type id"
// @Success 200 {object} model.OperationType
// @Failure 404 {object} model.ResponseMessage
// @Router /operation-type/{id} [get]
func (h *OperationTypeHandlers) GetOperationTypeById(c *gin.Context) {}

// @Summary update a operation type by ID
// @ID update-operation-type-by-id
// @Tags operation-types
// @Param id path string true "operation type id"
// @Param operation type body model.OperationType true "role"
// @Produce json
// @Consume json
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /operation-type/{id} [post]
func (h *OperationTypeHandlers) UpdateOperationTypeById(c *gin.Context) {}

// @Summary put a operation type by ID
// @ID update-operation-type-by-id
// @Tags operation-types
// @Param id path string true "operation type id"
// @Param operation type body model.OperationType true "role"
// @Produce json
// @Consume json
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /operation-type/{id} [put]
func (h *OperationTypeHandlers) PutOperationTypeById(c *gin.Context) {}

// @Summary delete a operation type by ID
// @ID delete-operation-type-by-id
// @Tags operation-types
// @Param id path string true "operation type id"
// @Produce json
// @Success 200 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /operation-type/{id} [delete]
func (h *OperationTypeHandlers) DeleteOperationTypeById(c *gin.Context) {}
