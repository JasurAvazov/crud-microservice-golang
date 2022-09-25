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

type categoryParams struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type updateCategoryParams struct {
	Name string `json:"name" binding:"required"`
}

// createCategory godoc swagger
// @Summary Create category
// @Description API to create a category
// @Router /category [POST]
// @Tags Category
// @Accept json
// @Produce json
// @Param request body categoryParams true "create category request parameters"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) createCategory(c *gin.Context) {
	var req categoryParams

	if err := c.ShouldBindJSON(&req); err != nil {

		h.log.Error(err.Error())
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation, err)

		return
	}
	model := models.Category{
		Id:   req.Id,
		Name: req.Name,
	}
	if err := h.ctrl.CreateCategory(c, model); err != nil {
		h.log.Error(err.Error())
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
		return
	}

	h.handleSuccessResponse(c, nil)
}

// readCategory godoc swagger
// @Summary Read one category
// @Description API to get a category by id
// @Router /category/{code} [GET]
// @Tags Category
// @Accept json
// @Produce json
// @Param code path string true "category id"
// @Success 200 {object} views.R{data=views.category}
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) readCategory(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation, errors.New("empty branch's ID"))
		return
	}
	r, err := h.ctrl.GetCategory(c, code)
	switch {
	case err == nil:
		h.handleSuccessResponse(c, views.Category(r))
	case errors.Is(err, errs.ErrNotFound):
		h.handleSuccessResponse(c, nil)
	default:
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
	}
}

// readAllCategories godoc swagger
// @Summary Read all categories
// @Description API to get all categories
// @Router /categories [GET]
// @Tags Category
// @Accept json
// @Produce json
// @Success 200 {object} views.R{data=[]views.category}
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) readAllCategories(c *gin.Context) {
	readCategory, err := h.ctrl.ListCategory(c)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
		return
	}
	h.handleSuccessResponse(c, views.Categories(readCategory))
}

// updateCategory godoc swagger
// @Summary Update category
// @Description API to update category by id
// @Router /category/{code} [PUT]
// @Tags Category
// @Accept json
// @Produce json
// @Param code path string true "category id"
// @Param request body updateCategoryParams true "update category request parameters"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) updateCategory(c *gin.Context) {
	var req updateCategoryParams
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
	model := models.Category{
		Id:   intId,
		Name: req.Name,
	}
	err = h.ctrl.UpdateCategory(c, model)
	switch {
	case err == nil:
		h.handleSuccessResponse(c, views.Category(model))
	case errors.Is(err, errs.ErrNotFound):
		h.handleErrorResponse(c, http.StatusNotFound, status.ErrorCodeDB, err)
	default:
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
	}
}

// deleteCategory godoc swagger
// @Summary Delete category
// @Description API to deactivate category by id
// @Router /category/{code} [DELETE]
// @Tags Category
// @Accept json
// @Produce json
// @Param code path string true "category id"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) deleteCategory(c *gin.Context) {
	id := c.Param("code")
	if id == "" {
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation,
			errors.New("empty branch ID"))
		return
	}
	err := h.ctrl.DeleteCategory(c, id)
	switch {
	case err == nil:
		h.handleSuccessResponse(c, nil)
	case errors.Is(err, errs.ErrNotFound):
		h.handleErrorResponse(c, http.StatusNotFound, status.ErrorCodeDB, err)
	default:
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
	}
}
