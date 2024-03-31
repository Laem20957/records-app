package rest

import (
	"net/http"
	"strconv"

	"github.com/Laem20957/records-app/internal/domain"
	"github.com/gin-gonic/gin"
)

// @Summary Create new record
// @Tags API
// @Accept json
// @Produce json
// @Param input body domain.UpdateRecord true "record info"
// @Success 200 {integer} integer 1
// @Failure 500,400,404 {object} domain.MessageResponse
// @Router /api/record [post]
func (h *Handler) create(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		domain.ServerResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	var input domain.Record
	if err := ctx.BindJSON(&input); err != nil {
		domain.ServerResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Services.IServiceRecordMethods.CreateRecords(ctx, userId, input)
	if err != nil {
		domain.ServerResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get record by id
// @Tags API
// @Accept json
// @Produce json
// @Param id path integer true "Record ID"
// @Success 200 {object} domain.Record
// @Failure 500,400,404 {object} domain.MessageResponse
// @Router /api/record/{id} [get]
func (h *Handler) getById(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		domain.ServerResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		domain.ServerResponse(ctx, http.StatusBadRequest, "invalid id param")
	}

	record, err := h.Services.GetByIDRecords(ctx, userId, id)
	if err != nil {
		domain.ServerResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, record)

}

// @Summary Get all records
// @Tags API
// @Accept json
// @Produce json
// @Success 200 {object} domain.GetAllRecordResponse
// @Failure 500,400,404 {object} domain.MessageResponse
// @Router /api/record [get]
func (h *Handler) getAll(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		domain.ServerResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	records, err := h.Services.IServiceRecordMethods.GetAllRecords(ctx, userId)
	if err != nil {
		domain.ServerResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, domain.GetAllRecordResponse{Data: records})
}

// @Summary Delete record by id
// @Tags API
// @Accept json
// @Produce json
// @Param id path integer true "Record ID"
// @Success 200 {integer} integer 1
// @Failure 500,400,404 {object} domain.MessageResponse
// @Router /api/record/{id} [delete]
func (h *Handler) delete(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		domain.ServerResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		domain.ServerResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.Services.DeleteRecords(ctx, userId, id)
	if err != nil {
		domain.ServerResponse(ctx, http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusOK, domain.StatusResponse{Status: "ok"})
}

// @Summary Update record by id
// @Tags API
// @Accept json
// @Produce json
// @Param id path integer true "Record ID"
// @Param input body domain.UpdateRecord true "record info"
// @Success 200 {integer} integer 1
// @Failure 500,400,404 {object} domain.MessageResponse
// @Router /api/record/{id} [put]
func (h *Handler) update(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		domain.ServerResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		domain.ServerResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	var input domain.UpdateRecord
	if err := ctx.BindJSON(&input); err != nil {
		domain.ServerResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.Services.UpdateRecords(ctx, userId, id, input); err != nil {
		domain.ServerResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, domain.StatusResponse{Status: "ok"})
}
