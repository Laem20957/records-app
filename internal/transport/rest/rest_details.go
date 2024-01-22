package rest

import (
	"net/http"
	"strconv"

	domain "github.com/Laem20957/records-app/internal/domains"
	"github.com/gin-gonic/gin"
)

// @Summary Create new record
// @Security ApiKeyAuth
// @Tags record
// @Description Create record
// @ID Create-record
// @Accept json
// @Produce json
// @Param input body domain.UpdateNote true "note info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /api/note [post]

func (h *Handler) create(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		domain.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input domain.Record
	if err := c.BindJSON(&input); err != nil {
		domain.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Services.ServiceRecordMethods.CreateRecords(c, userId, input)
	if err != nil {
		domain.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		domain.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		domain.NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
	}

	note, err := h.Services.GetByIDRecords(c, userId, id)
	if err != nil {
		domain.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, note)

}

func (h *Handler) getAll(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		domain.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	notes, err := h.Services.ServiceRecordMethods.GetAllRecords(c, userId)
	if err != nil {
		domain.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, domain.GetAllRecordResponse{Data: notes})
}

func (h *Handler) delete(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		domain.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		domain.NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.Services.DeleteRecords(c, userId, id)
	if err != nil {
		domain.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, domain.StatusResponse{Status: "ok"})
}

func (h *Handler) update(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		domain.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		domain.NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input domain.UpdateRecord
	if err := c.BindJSON(&input); err != nil {
		domain.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.Services.UpdateRecords(c, userId, id, input); err != nil {
		domain.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, domain.StatusResponse{Status: "ok"})
}
