package rest

import (
	"net/http"
	"strconv"

	"github.com/Laem20957/records-app/internal/domain"
	"github.com/gin-gonic/gin"
)

// note.go

// @Summary Create new record
// @Security ApiKeyAuth
// @Tags record
// @Description Create record
// @ID Create-record
// @Accept json
// @Produce json
// @Param input body domain.UpdateRecord true "record info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} domain.MessageResponse
// @Failure 500 {object} domain.MessageResponse
// @Failure default {object} domain.MessageResponse
// @Router /api/records [post]

func (h *Handler) create(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		domain.ServerResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input domain.Record
	if err := c.BindJSON(&input); err != nil {
		domain.ServerResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Services.IServiceRecordMethods.CreateRecords(c, userId, input)
	if err != nil {
		domain.ServerResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get record by id
// @Security ApiKeyAuth
// @Tags record
// @Description Get record by id
// @ID Get-record-by-id
// @Accept json
// @Produce json
// @Param id path integer true "Record ID"
// @Success 200 {object} domain.Record
// @Failure 400,404 {object} domain.MessageResponse
// @Failure 500 {object} domain.MessageResponse
// @Failure default {object} domain.MessageResponse
// @Router /api/record/{id} [get]

func (h *Handler) getById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		domain.ServerResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		domain.ServerResponse(c, http.StatusBadRequest, "invalid id param")
	}

	note, err := h.Services.GetByIDRecords(c, userId, id)
	if err != nil {
		domain.ServerResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, note)

}

// @Summary Get all records
// @Security ApiKeyAuth
// @Tags record
// @Description Get all records
// @ID Get-all-record
// @Accept json
// @Produce json
// @Success 200 {object} domain.GetAllRecordResponse
// @Failure 400,404 {object} domain.MessageResponse
// @Failure 500 {object} domain.MessageResponse
// @Failure default {object} domain.MessageResponse
// @Router /api/record [get]

func (h *Handler) getAll(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		domain.ServerResponse(c, http.StatusInternalServerError, err.Error())
	}

	notes, err := h.Services.IServiceRecordMethods.GetAllRecords(c, userId)
	if err != nil {
		domain.ServerResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, domain.GetAllRecordResponse{Data: notes})
}

// @Summary Delete record by id
// @Security ApiKeyAuth
// @Tags record
// @Description Delete record by id
// @ID Delete-record-by-id
// @Accept json
// @Produce json
// @Param id path integer true "Record ID"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} domain.MessageResponse
// @Failure 500 {object} domain.MessageResponse
// @Failure default {object} domain.MessageResponse
// @Router /api/record/{id} [delete]

func (h *Handler) delete(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		domain.ServerResponse(c, http.StatusInternalServerError, err.Error())
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		domain.ServerResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.Services.DeleteRecords(c, userId, id)
	if err != nil {
		domain.ServerResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, domain.StatusResponse{Status: "ok"})
}

// @Summary Update record by id
// @Security ApiKeyAuth
// @Tags record
// @Description Update record by id
// @ID Update-record-by-id
// @Accept json
// @Produce json
// @Param id path integer true "Record ID"
// @Param input body domain.UpdateRecord true "record info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} domain.MessageResponse
// @Failure 500 {object} domain.MessageResponse
// @Failure default {object} domain.MessageResponse
// @Router /api/record/{id} [put]

func (h *Handler) update(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		domain.ServerResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		domain.ServerResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input domain.UpdateRecord
	if err := c.BindJSON(&input); err != nil {
		domain.ServerResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.Services.UpdateRecords(c, userId, id, input); err != nil {
		domain.ServerResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, domain.StatusResponse{Status: "ok"})
}
