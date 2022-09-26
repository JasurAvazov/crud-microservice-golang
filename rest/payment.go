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
	"time"
)

type paymentParams struct {
	Id     int       `json:"id" binding:"required"`
	Time   time.Time `json:"time" binding:"required"`
	Amount float64   `json:"amount" binding:"required"`
	InvId  int       `json:"inv_id" binding:"required"`
}

type updatePaymentParams struct {
	Time   time.Time `json:"time" binding:"required"`
	Amount float64   `json:"amount" binding:"required"`
	InvId  int       `json:"inv_id" binding:"required"`
}

// createPayment godoc swagger
// @Summary Create payment
// @Description API to create a payment
// @Router /payment [POST]
// @Tags Payment
// @Accept json
// @Produce json
// @Param request body paymentParams true "create payment request parameters"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) createPayment(c *gin.Context) {
	var req paymentParams

	if err := c.ShouldBindJSON(&req); err != nil {

		h.log.Error(err.Error())
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation, err)

		return
	}
	model := models.Payment{
		Id:     req.Id,
		Time:   req.Time,
		Amount: req.Amount,
		InvId:  req.InvId,
	}
	if err := h.ctrl.CreatePayment(c, model); err != nil {
		h.log.Error(err.Error())
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
		return
	}

	h.handleSuccessResponse(c, nil)
}

// readPayment godoc swagger
// @Summary Read one payment
// @Description API to get a payment by id
// @Router /payment/{code} [GET]
// @Tags Payment
// @Accept json
// @Produce json
// @Param code path string true "payment id"
// @Success 200 {object} views.R{data=views.payment}
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) readPayment(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation, errors.New("empty branch's ID"))
		return
	}
	r, err := h.ctrl.GetPayment(c, code)
	switch {
	case err == nil:
		h.handleSuccessResponse(c, views.Payment(r))
	case errors.Is(err, errs.ErrNotFound):
		h.handleSuccessResponse(c, nil)
	default:
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
	}
}

// readAllPayments godoc swagger
// @Summary Read all payments
// @Description API to get all payments
// @Router /payments [GET]
// @Tags Payment
// @Accept json
// @Produce json
// @Success 200 {object} views.R{data=[]views.payment}
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) readAllPayments(c *gin.Context) {
	readPayment, err := h.ctrl.ListPayment(c)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
		return
	}
	h.handleSuccessResponse(c, views.Payments(readPayment))
}

// updatePayment godoc swagger
// @Summary Update payment
// @Description API to update payment by id
// @Router /payment/{code} [PUT]
// @Tags Payment
// @Accept json
// @Produce json
// @Param code path string true "payment id"
// @Param request body updatePaymentParams true "update payment request parameters"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) updatePayment(c *gin.Context) {
	var req updatePaymentParams
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
	model := models.Payment{
		Id:     intId,
		Time:   req.Time,
		Amount: req.Amount,
		InvId:  req.InvId,
	}
	err = h.ctrl.UpdatePayment(c, model)
	switch {
	case err == nil:
		h.handleSuccessResponse(c, views.Payment(model))
	case errors.Is(err, errs.ErrNotFound):
		h.handleErrorResponse(c, http.StatusNotFound, status.ErrorCodeDB, err)
	default:
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
	}
}

// deletePayment godoc swagger
// @Summary Delete payment
// @Description API to deactivate payment by id
// @Router /payment/{code} [DELETE]
// @Tags Payment
// @Accept json
// @Produce json
// @Param code path string true "payment id"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) deletePayment(c *gin.Context) {
	id := c.Param("code")
	if id == "" {
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation,
			errors.New("empty branch ID"))
		return
	}
	err := h.ctrl.DeletePayment(c, id)
	switch {
	case err == nil:
		h.handleSuccessResponse(c, nil)
	case errors.Is(err, errs.ErrNotFound):
		h.handleErrorResponse(c, http.StatusNotFound, status.ErrorCodeDB, err)
	default:
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
	}
}
