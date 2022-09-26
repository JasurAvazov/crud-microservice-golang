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

type invoiceParams struct {
	Id     int       `json:"id" binding:"required"`
	OrdId  int       `json:"ord_id" binding:"required"`
	Amount float64   `json:"amount" binding:"required"`
	Issued time.Time `json:"issued" binding:"required"`
	Due    time.Time `json:"due" binding:"required"`
}

type updateInvoiceParams struct {
	OrdId  int       `json:"ord_id" binding:"required"`
	Amount float64   `json:"amount" binding:"required"`
	Issued time.Time `json:"issued" binding:"required"`
	Due    time.Time `json:"due" binding:"required"`
}

// createInvoice godoc swagger
// @Summary Create invoice
// @Description API to create an invoice
// @Router /invoice [POST]
// @Tags Invoice
// @Accept json
// @Produce json
// @Param request body invoiceParams true "create invoice request parameters"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) createInvoice(c *gin.Context) {
	var req invoiceParams

	if err := c.ShouldBindJSON(&req); err != nil {

		h.log.Error(err.Error())
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation, err)

		return
	}
	model := models.Invoice{
		Id:     req.Id,
		OrdId:  req.OrdId,
		Amount: req.Amount,
		Issued: req.Issued,
		Due:    req.Due,
	}
	if err := h.ctrl.CreateInvoice(c, model); err != nil {
		h.log.Error(err.Error())
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
		return
	}

	h.handleSuccessResponse(c, nil)
}

// readInvoice godoc swagger
// @Summary Read one invoice
// @Description API to get a invoice by id
// @Router /invoice/{code} [GET]
// @Tags Invoice
// @Accept json
// @Produce json
// @Param code path string true "invoice id"
// @Success 200 {object} views.R{data=views.invoice}
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) readInvoice(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation, errors.New("empty branch's ID"))
		return
	}
	r, err := h.ctrl.GetInvoice(c, code)
	switch {
	case err == nil:
		h.handleSuccessResponse(c, views.Invoice(r))
	case errors.Is(err, errs.ErrNotFound):
		h.handleSuccessResponse(c, nil)
	default:
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
	}
}

// readAllInvoices godoc swagger
// @Summary Read all invoices
// @Description API to get all invoices
// @Router /invoices [GET]
// @Tags Invoice
// @Accept json
// @Produce json
// @Success 200 {object} views.R{data=[]views.invoice}
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) readAllInvoices(c *gin.Context) {
	readInvoice, err := h.ctrl.ListInvoice(c)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
		return
	}
	h.handleSuccessResponse(c, views.Invoices(readInvoice))
}

// updateInvoice godoc swagger
// @Summary Update invoice
// @Description API to update invoice by id
// @Router /invoice/{code} [PUT]
// @Tags Invoice
// @Accept json
// @Produce json
// @Param code path string true "invoice id"
// @Param request body updateInvoiceParams true "update invoice request parameters"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) updateInvoice(c *gin.Context) {
	var req updateInvoiceParams
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
	model := models.Invoice{
		Id:     intId,
		OrdId:  req.OrdId,
		Amount: req.Amount,
		Issued: req.Issued,
		Due:    req.Due,
	}
	err = h.ctrl.UpdateInvoice(c, model)
	switch {
	case err == nil:
		h.handleSuccessResponse(c, views.Invoice(model))
	case errors.Is(err, errs.ErrNotFound):
		h.handleErrorResponse(c, http.StatusNotFound, status.ErrorCodeDB, err)
	default:
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
	}
}

// deleteInvoice godoc swagger
// @Summary Delete invoice
// @Description API to deactivate invoice by id
// @Router /invoice/{code} [DELETE]
// @Tags Invoice
// @Accept json
// @Produce json
// @Param code path string true "invoice id"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) deleteInvoice(c *gin.Context) {
	id := c.Param("code")
	if id == "" {
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation,
			errors.New("empty branch ID"))
		return
	}
	err := h.ctrl.DeleteInvoice(c, id)
	switch {
	case err == nil:
		h.handleSuccessResponse(c, nil)
	case errors.Is(err, errs.ErrNotFound):
		h.handleErrorResponse(c, http.StatusNotFound, status.ErrorCodeDB, err)
	default:
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
	}
}
