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

type orderParams struct {
	Id      int       `json:"id" binding:"required"`
	Date    time.Time `json:"date" binding:"required"`
	Cust_id int       `json:"cust_id" binding:"required"`
}

type updateOrderParams struct {
	Date    time.Time `json:"date" binding:"required"`
	Cust_id int       `json:"cust_id" binding:"required"`
}

// createOrder godoc swagger
// @Summary Create order
// @Description API to create a order
// @Router /order [POST]
// @Tags Order
// @Accept json
// @Produce json
// @Param request body orderParams true "create order request parameters"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) createOrder(c *gin.Context) {
	var req orderParams

	if err := c.ShouldBindJSON(&req); err != nil {

		h.log.Error(err.Error())
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation, err)

		return
	}
	model := models.Order{
		Id:      req.Id,
		Date:    req.Date,
		Cust_id: req.Cust_id,
	}
	if err := h.ctrl.CreateOrder(c, model); err != nil {
		h.log.Error(err.Error())
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
		return
	}

	h.handleSuccessResponse(c, nil)
}

// readOrder godoc swagger
// @Summary Read one order
// @Description API to get a order by id
// @Router /order/{code} [GET]
// @Tags Order
// @Accept json
// @Produce json
// @Param code path string true "order id"
// @Success 200 {object} views.R{data=views.order}
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) readOrder(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation, errors.New("empty branch's ID"))
		return
	}
	r, err := h.ctrl.GetOrder(c, code)
	switch {
	case err == nil:
		h.handleSuccessResponse(c, views.Order(r))
	case errors.Is(err, errs.ErrNotFound):
		h.handleSuccessResponse(c, nil)
	default:
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
	}
}

// readAllOrders godoc swagger
// @Summary Read all orders
// @Description API to get all orders
// @Router /orders [GET]
// @Tags Order
// @Accept json
// @Produce json
// @Success 200 {object} views.R{data=[]views.order}
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) readAllOrders(c *gin.Context) {
	readOrder, err := h.ctrl.ListOrder(c)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
		return
	}
	h.handleSuccessResponse(c, views.Orders(readOrder))
}

// updateOrder godoc swagger
// @Summary Update order
// @Description API to update order by id
// @Router /order/{code} [PUT]
// @Tags Order
// @Accept json
// @Produce json
// @Param code path string true "order id"
// @Param request body updateOrderParams true "update order request parameters"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) updateOrder(c *gin.Context) {
	var req updateOrderParams
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
	model := models.Order{
		Id:      intId,
		Date:    req.Date,
		Cust_id: req.Cust_id,
	}
	err = h.ctrl.UpdateOrder(c, model)
	switch {
	case err == nil:
		h.handleSuccessResponse(c, views.Order(model))
	case errors.Is(err, errs.ErrNotFound):
		h.handleErrorResponse(c, http.StatusNotFound, status.ErrorCodeDB, err)
	default:
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
	}
}

// deleteOrder godoc swagger
// @Summary Delete order
// @Description API to deactivate order by id
// @Router /order/{code} [DELETE]
// @Tags Order
// @Accept json
// @Produce json
// @Param code path string true "order id"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) deleteOrder(c *gin.Context) {
	id := c.Param("code")
	if id == "" {
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation,
			errors.New("empty branch ID"))
		return
	}
	err := h.ctrl.DeleteOrder(c, id)
	switch {
	case err == nil:
		h.handleSuccessResponse(c, nil)
	case errors.Is(err, errs.ErrNotFound):
		h.handleErrorResponse(c, http.StatusNotFound, status.ErrorCodeDB, err)
	default:
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
	}
}
