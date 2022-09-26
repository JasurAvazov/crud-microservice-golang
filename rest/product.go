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

type productParams struct {
	Id          int     `json:"id" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	CategoryId  int     `json:"category_id" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Photo       string  `json:"photo" binding:"required"`
}

type updateProductParams struct {
	Name        string  `json:"name" binding:"required"`
	CategoryId  int     `json:"category_id" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Photo       string  `json:"photo" binding:"required"`
}

// createProduct godoc swagger
// @Summary Create product
// @Description API to create a product
// @Router /product [POST]
// @Tags Product
// @Accept json
// @Produce json
// @Param request body productParams true "create product request parameters"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) createProduct(c *gin.Context) {
	var req productParams

	if err := c.ShouldBindJSON(&req); err != nil {

		h.log.Error(err.Error())
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation, err)

		return
	}
	model := models.Product{
		Id:          req.Id,
		Name:        req.Name,
		CategoryId:  req.CategoryId,
		Description: req.Description,
		Price:       req.Price,
		Photo:       req.Photo,
	}
	if err := h.ctrl.CreateProduct(c, model); err != nil {
		h.log.Error(err.Error())
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
		return
	}

	h.handleSuccessResponse(c, nil)
}

// readProduct godoc swagger
// @Summary Read one product
// @Description API to get a product by id
// @Router /product/{code} [GET]
// @Tags Product
// @Accept json
// @Produce json
// @Param code path string true "product id"
// @Success 200 {object} views.R{data=views.product}
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) readProduct(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation, errors.New("empty branch's ID"))
		return
	}
	r, err := h.ctrl.GetProduct(c, code)
	switch {
	case err == nil:
		h.handleSuccessResponse(c, views.Product(r))
	case errors.Is(err, errs.ErrNotFound):
		h.handleSuccessResponse(c, nil)
	default:
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
	}
}

// readAllProducts godoc swagger
// @Summary Read all products
// @Description API to get all products
// @Router /products [GET]
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {object} views.R{data=[]views.product}
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) readAllProducts(c *gin.Context) {
	readProduct, err := h.ctrl.ListProduct(c)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
		return
	}
	h.handleSuccessResponse(c, views.Products(readProduct))
}

// updateProduct godoc swagger
// @Summary Update product
// @Description API to update product by id
// @Router /product/{code} [PUT]
// @Tags Product
// @Accept json
// @Produce json
// @Param code path string true "product id"
// @Param request body updateProductParams true "update product request parameters"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) updateProduct(c *gin.Context) {
	var req updateProductParams
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
	model := models.Product{
		Id:          intId,
		Name:        req.Name,
		CategoryId:  req.CategoryId,
		Description: req.Description,
		Price:       req.Price,
		Photo:       req.Photo,
	}
	err = h.ctrl.UpdateProduct(c, model)
	switch {
	case err == nil:
		h.handleSuccessResponse(c, views.Product(model))
	case errors.Is(err, errs.ErrNotFound):
		h.handleErrorResponse(c, http.StatusNotFound, status.ErrorCodeDB, err)
	default:
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
	}
}

// deleteProduct godoc swagger
// @Summary Delete product
// @Description API to deactivate product by id
// @Router /product/{code} [DELETE]
// @Tags Product
// @Accept json
// @Produce json
// @Param code path string true "product id"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) deleteProduct(c *gin.Context) {
	id := c.Param("code")
	if id == "" {
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation,
			errors.New("empty branch ID"))
		return
	}
	err := h.ctrl.DeleteProduct(c, id)
	switch {
	case err == nil:
		h.handleSuccessResponse(c, nil)
	case errors.Is(err, errs.ErrNotFound):
		h.handleErrorResponse(c, http.StatusNotFound, status.ErrorCodeDB, err)
	default:
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
	}
}
