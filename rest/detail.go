package rest

import (
	"apelsin/errs"
	"apelsin/models"
	"apelsin/pkg/status"
	"apelsin/rest/views"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type detailParams struct {
	Id       int   `json:"id" binding:"required"`
	OrdId    int   `json:"ord_id" binding:"required"`
	PrId     int   `json:"pr_id" binding:"required"`
	Quantity int16 `json:"quantity" binding:"required"`
}

type updateDetailParams struct {
	OrdId    int   `json:"ord_id" binding:"required"`
	PrId     int   `json:"pr_id" binding:"required"`
	Quantity int16 `json:"quantity" binding:"required"`
}

// createDetail godoc swagger
// @Summary Create detail
// @Description API to create a detail
// @Router /detail [POST]
// @Tags Detail
// @Accept json
// @Produce json
// @Param request body detailParams true "create detail request parameters"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) createDetail(c *gin.Context) {
	var req detailParams

	if err := c.ShouldBindJSON(&req); err != nil {

		h.log.Error(err.Error())
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation, err)

		return
	}
	model := models.Detail{
		Id:       req.Id,
		OrdId:    req.OrdId,
		PrId:     req.PrId,
		Quantity: req.Quantity,
	}
	if err := h.ctrl.CreateDetail(c, model); err != nil {
		h.log.Error(err.Error())
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
		return
	}

	h.handleSuccessResponse(c, nil)
}

// readDetail godoc swagger
// @Summary Read one detail
// @Description API to get a detail by id
// @Router /detail/{code} [GET]
// @Tags Detail
// @Accept json
// @Produce json
// @Param code path string true "detail id"
// @Success 200 {object} views.R{data=views.detail}
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) readDetail(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation, errors.New("empty branch's ID"))
		return
	}
	r, err := h.ctrl.GetDetail(c, code)
	switch {
	case err == nil:
		h.handleSuccessResponse(c, views.Detail(r))
	case errors.Is(err, errs.ErrNotFound):
		h.handleSuccessResponse(c, nil)
	default:
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
	}
}

// readAllDetails godoc swagger
// @Summary Read all details
// @Description API to get all details
// @Router /details [GET]
// @Tags Detail
// @Accept json
// @Produce json
// @Success 200 {object} views.R{data=[]views.detail}
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) readAllDetails(c *gin.Context) {
	readDetail, err := h.ctrl.ListDetail(c)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
		return
	}
	h.handleSuccessResponse(c, views.Details(readDetail))
}

// updateDetail godoc swagger
// @Summary Update detail
// @Description API to update detail by id
// @Router /detail/{code} [PUT]
// @Tags Detail
// @Accept json
// @Produce json
// @Param code path string true "detail id"
// @Param request body updateDetailParams true "update detail request parameters"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) updateDetail(c *gin.Context) {
	var req updateDetailParams
	code := c.Param("code")
	intId, err := strconv.Atoi(code)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, status.ErrorConvert, err)
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error(err.Error())

		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation, err)
		return
	}
	model := models.Detail{
		Id:       intId,
		OrdId:    req.OrdId,
		PrId:     req.PrId,
		Quantity: req.Quantity,
	}
	err = h.ctrl.UpdateDetail(c, model)
	switch {
	case err == nil:
		h.handleSuccessResponse(c, views.Detail(model))
	case errors.Is(err, errs.ErrNotFound):
		h.handleErrorResponse(c, http.StatusNotFound, status.ErrorCodeDB, err)
	default:
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
	}
}

// deleteDetail godoc swagger
// @Summary Delete detail
// @Description API to deactivate detail by id
// @Router /detail/{code} [DELETE]
// @Tags Detail
// @Accept json
// @Produce json
// @Param code path string true "detail id"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) deleteDetail(c *gin.Context) {
	id := c.Param("code")
	if id == "" {
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation,
			errors.New("empty branch ID"))
		return
	}
	err := h.ctrl.DeleteDetail(c, id)
	switch {
	case err == nil:
		h.handleSuccessResponse(c, nil)
	case errors.Is(err, errs.ErrNotFound):
		h.handleErrorResponse(c, http.StatusNotFound, status.ErrorCodeDB, err)
	default:
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
	}
}
