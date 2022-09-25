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

type customerParams struct {
	Id      int    `json:"id" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Country string `json:"country" binding:"required"`
	Address string `json:"address" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
}

type updateCustomerParams struct {
	Name    string `json:"name" binding:"required"`
	Country string `json:"country" binding:"required"`
	Address string `json:"address" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
}

// createCustomer godoc swagger
// @Summary Create customer
// @Description API to create a customer
// @Router /customer [POST]
// @Tags Customer
// @Accept json
// @Produce json
// @Param request body customerParams true "create customer request parameters"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) createCustomer(c *gin.Context) {
	var req customerParams

	if err := c.ShouldBindJSON(&req); err != nil {

		h.log.Error(err.Error())
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation, err)

		return
	}
	model := models.Customer{
		Id:      req.Id,
		Name:    req.Name,
		Country: req.Country,
		Address: req.Address,
		Phone:   req.Phone,
	}
	if err := h.ctrl.CreateCustomer(c, model); err != nil {
		h.log.Error(err.Error())
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
		return
	}

	h.handleSuccessResponse(c, nil)
}

// readCustomer godoc swagger
// @Summary Read one customer
// @Description API to get a customer by id
// @Router /customer/{code} [GET]
// @Tags Customer
// @Accept json
// @Produce json
// @Param code path string true "customer id"
// @Success 200 {object} views.R{data=views.customer}
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) readCustomer(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation, errors.New("empty branch's ID"))
		return
	}
	r, err := h.ctrl.GetCustomer(c, code)
	switch {
	case err == nil:
		h.handleSuccessResponse(c, views.Customer(r))
	case errors.Is(err, errs.ErrNotFound):
		h.handleSuccessResponse(c, nil)
	default:
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
	}
}

// readAllCustomers godoc swagger
// @Summary Read all customers
// @Description API to get all customers
// @Router /customers [GET]
// @Tags Customer
// @Accept json
// @Produce json
// @Success 200 {object} views.R{data=[]views.customer}
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) readAllCustomers(c *gin.Context) {
	readCustomer, err := h.ctrl.ListCustomer(c)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
		return
	}
	h.handleSuccessResponse(c, views.Customers(readCustomer))
}

// updateCustomer godoc swagger
// @Summary Update customer
// @Description API to update customer by id
// @Router /customer/{code} [PUT]
// @Tags Customer
// @Accept json
// @Produce json
// @Param code path string true "customer id"
// @Param request body updateCustomerParams true "update customer request parameters"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) updateCustomer(c *gin.Context) {
	var req updateCustomerParams
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
	model := models.Customer{
		Id:      intId,
		Name:    req.Name,
		Country: req.Country,
		Address: req.Address,
		Phone:   req.Phone,
	}
	err = h.ctrl.UpdateCustomer(c, model)
	switch {
	case err == nil:
		h.handleSuccessResponse(c, views.Customer(model))
	case errors.Is(err, errs.ErrNotFound):
		h.handleErrorResponse(c, http.StatusNotFound, status.ErrorCodeDB, err)
	default:
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
	}
}

// deleteCustomer godoc swagger
// @Summary Delete customer
// @Description API to deactivate customer by id
// @Router /customer/{code} [DELETE]
// @Tags Customer
// @Accept json
// @Produce json
// @Param code path string true "customer id"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) deleteCustomer(c *gin.Context) {
	id := c.Param("code")
	if id == "" {
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation,
			errors.New("empty branch ID"))
		return
	}
	err := h.ctrl.DeleteCustomer(c, id)
	switch {
	case err == nil:
		h.handleSuccessResponse(c, nil)
	case errors.Is(err, errs.ErrNotFound):
		h.handleErrorResponse(c, http.StatusNotFound, status.ErrorCodeDB, err)
	default:
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
	}
}
