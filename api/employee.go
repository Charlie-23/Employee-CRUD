package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/pzanwar/employee/db/sqlc"
)

type createEmployeeRequest struct {
	Name     string  `json:"name" binding:"required"`
	Salary   float64 `json:"salary" binding:"required"`
	Position string  `json:"position" binding:"required"`
}

func (server *Server) createEmployee(ctx *gin.Context) {
	var req createEmployeeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateEmployeeParams{
		Name:     req.Name,
		Salary:   req.Salary,
		Position: req.Position,
	}

	account, err := server.store.CreateEmployee(ctx, arg)
	if err != nil {
		errCode := db.ErrorCode(err)
		if errCode == db.ForeignKeyViolation || errCode == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type getEmployeeRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getEmployeeByID(ctx *gin.Context) {
	var req getEmployeeRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	employee, err := server.store.GetEmployeeByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, employee)
}

type updateEmployeeRequest struct {
	Name     string  `json:"name"`
	Salary   float64 `json:"salary"`
	Position string  `json:"position"`
}

func (server *Server) updateEmployee(ctx *gin.Context) {
	var req updateEmployeeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var reqid getEmployeeRequest
	if err := ctx.ShouldBindUri(&reqid); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	employee, err := server.store.GetEmployeeByID(ctx, reqid.ID)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if req.Name == "" {
		req.Name = employee.Name
	}
	if req.Position == "" {
		req.Position = employee.Position
	}
	if req.Salary == 0 {
		req.Salary = employee.Salary
	}

	arg := db.UpdateEmployeeParams{
		ID:       reqid.ID,
		Salary:   req.Salary,
		Name:     req.Name,
		Position: req.Position,
	}

	result, err := server.store.UpdateEmployee(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, result)
}

type deleteEmployeeRequest struct {
	ID int64 `uri:"id" binding:"required"`
}
type deleteEmployeeResponse struct {
	ID int64
	Status string
}

func (server *Server) deleteEmployee(ctx *gin.Context) {
	var req deleteEmployeeRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteEmployee(ctx, req.ID)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	res := deleteEmployeeResponse{
		ID: req.ID,
		Status: "DELETED",
	}
	ctx.JSON(http.StatusOK, res)
}
