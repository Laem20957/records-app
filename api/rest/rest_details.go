package rest

import (
	"net/http"
	"strconv"

	handler "records-app/api/rest/v1/handlers"
	domain "records-app/internal/domain"

	"github.com/gin-gonic/gin"
)

// @Summary Get all records
// @Tags API
// @Accept json
// @Produce json
// @Success 200 {object} domain.GetAllRecordResponse
// @Failure 500,400,404 {object} domain.MessageResponse
// @Router /api/all/record [get]
func GetAll(ctx *gin.Context) {
	ctx, err := handler.GetAllContext(ctx)
	if err != nil {
		domain.ServerResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	records, err := handler.Handler{}.Services.IServiceRecordMethods.GetAllRecords(ctx)
	if err != nil {
		domain.ServerResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, domain.GetAllRecordResponse{Data: records})
}

// @Summary Get record by id
// @Tags API
// @Accept json
// @Produce json
// @Param id path integer true "Record ID"
// @Success 200 {object} domain.Records
// @Failure 500,400,404 {object} domain.MessageResponse
// @Router /api/record/{id} [get]
func GetById(ctx *gin.Context) {
	ctx, err := handler.GetUserContext(ctx)
	if err != nil {
		domain.ServerResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	recordId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		domain.ServerResponse(ctx, http.StatusBadRequest, "invalid recordId was passed")
		return
	}
	record, err := handler.Handler{}.Services.IServiceRecordMethods.GetByIDRecords(ctx, recordId)
	if err != nil {
		domain.ServerResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, record)
}

// @Summary Create new record
// @Tags API
// @Accept json
// @Produce json
// @Param input body domain.Records true "record info"
// @Success 200 {integer} integer 200
// @Failure 500,400,404 {object} domain.MessageResponse
// @Router /api/new/record [post]
func Create(ctx *gin.Context) {
	var input domain.Records
	ctx, err := handler.GetUserContext(ctx)
	if err != nil {
		domain.ServerResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	if err := ctx.BindJSON(&input); err != nil {
		domain.ServerResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	id, err := handler.Handler{}.Services.IServiceRecordMethods.CreateRecords(ctx)
	if err != nil {
		domain.ServerResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Update record by id
// @Tags API
// @Accept json
// @Produce json
// @Param id path integer true "Record ID"
// @Param input body domain.Records true "record info"
// @Success 200 {integer} integer 200
// @Failure 500,400,404 {object} domain.MessageResponse
// @Router /api/record/{id} [put]
func Update(ctx *gin.Context) {
	var record domain.Records
	ctx, err := handler.GetUserContext(ctx)
	if err != nil {
		domain.ServerResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	if err := ctx.BindJSON(&record); err != nil {
		domain.ServerResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	recordId, err := handler.Handler{}.Services.IServiceRecordMethods.UpdateRecords(ctx, record.ID, record.Title, record.Description)
	if err != nil {
		domain.ServerResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, domain.ResultResponse{Status: "OK", ResponseId: recordId.ID})
}

// @Summary Delete record by id
// @Tags API
// @Accept json
// @Produce json
// @Param id path integer true "Record ID"
// @Success 200 {integer} integer 200
// @Failure 500,400,404 {object} domain.MessageResponse
// @Router /api/record/{id} [delete]
func Delete(ctx *gin.Context) {
	ctx, err := handler.GetUserContext(ctx)
	if err != nil {
		domain.ServerResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	recordId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		domain.ServerResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}
	recordId, err = handler.Handler{}.Services.IServiceRecordMethods.DeleteRecords(ctx, recordId)
	if err != nil {
		domain.ServerResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, domain.ResultResponse{Status: "OK", ResponseId: recordId})
}
