package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type OperationLogHandlers struct {
	lg logger.Logger
}

func NewOperationLog(lg logger.Logger) *OperationLogHandlers {
	return &OperationLogHandlers{lg: lg}
}

// AddOperationLog
// @Summary add new item to the operation logs list
// @ID create-operation-log
// @Tags operation-logs
// @Produce json
// @Consume json
// @Param operation-log body model.OperationLog true "operation-log"
// @Success 201 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /operation-log [post]
func (h *OperationLogHandlers) AddOperationLog(c *gin.Context) {}

// GetOperationLogs
// @Summary get all items in the operation log list
// @ID get-all-operation-logs
// @Tags operation-logs
// @Produce json
// @Success 200 {object} []model.OperationLog
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /operation-log [get]
func (h *OperationLogHandlers) GetOperationLogs(c *gin.Context) {}

// GetOperationLogById
// @Summary get a operation log by ID
// @ID get-operation-log-by-id
// @Tags operation-logs
// @Produce json
// @Param id path string true "operation-log id"
// @Success 200 {object} model.OperationLog
// @Failure 404 {object} model.ResponseMessage
// @Router /operation-log/{id} [get]
func (h *OperationLogHandlers) GetOperationLogById(c *gin.Context) {}

// UpdateOperationLogById
// @Summary update a operation log by ID
// @ID update-operation-log-by-id
// @Tags operation-logs
// @Param id path string true "operation log id"
// @Param operation_log body model.OperationLog true "role"
// @Produce json
// @Consume json
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /operation-log/{id} [post]
func (h *OperationLogHandlers) UpdateOperationLogById(c *gin.Context) {}

// PutOperationLogById
// @Summary put a operation log by ID
// @ID update-operation-log-by-id
// @Tags operation-logs
// @Param id path string true "operation log id"
// @Param operation_log body model.OperationLog true "role"
// @Produce json
// @Consume json
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /operation-log/{id} [put]
func (h *OperationLogHandlers) PutOperationLogById(c *gin.Context) {}

// DeleteOperationLogById
// @Summary delete a operation log by ID
// @ID delete-operation-log-by-id
// @Tags operation-logs
// @Param id path string true "operation log id"
// @Produce json
// @Success 200 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /operation-log/{id} [delete]
func (h *OperationLogHandlers) DeleteOperationLogById(c *gin.Context) {}
