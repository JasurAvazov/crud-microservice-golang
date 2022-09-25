package rest

import (
	"apelsin/errs"
	"apelsin/models"
	"apelsin/pkg/status"
	"apelsin/rest/views"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type districtParams struct {
	CodeSOATO    string    `json:"code_soato" binding:"required"`
	Code         string    `json:"code" binding:"required"`
	CodeGNI      string    `json:"code_gni" binding:"required"`
	CodeProvince string    `json:"code_province" binding:"required"`
	Title        Languages `json:"title" binding:"required"`
	IsActive     bool      `json:"is_active" binding:"required"`
}

type updateDistrictParams struct {
	CodeSOATO    string    `json:"code_soato" binding:"required"`
	CodeGNI      string    `json:"code_gni" binding:"required"`
	CodeProvince string    `json:"code_province" binding:"required"`
	Title        Languages `json:"title" binding:"required"`
	IsActive     bool      `json:"is_active" binding:"required"`
}

// createDistrict godoc swagger
// @Summary Create district
// @Description API to create a district
// @Router /districts [POST]
// @Tags District
// @Accept json
// @Produce json
// @Param request body districtParams true "create district request parameters"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) createDistrict(c *gin.Context) {
	var req districtParams

	if err := c.ShouldBindJSON(&req); err != nil {

		h.log.Error(err.Error())
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation, err)

		return
	}
	model := models.District{
		CodeSOATO:    req.CodeSOATO,
		Code:         req.Code,
		CodeGNI:      req.CodeGNI,
		CodeProvince: req.CodeProvince,
		Title: models.Languages{
			Uz: req.Title.Uz,
			Ru: req.Title.Ru,
			En: req.Title.En,
		},
		ActivatedAt: nil,
		State:       false,
	}
	if req.IsActive {
		now := time.Now()
		model.ActivatedAt = &now
	}
	if err := h.ctrl.CreateDistrict(c, model); err != nil {
		h.log.Error(err.Error())
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
		return
	}

	h.handleSuccessResponse(c, nil)
}

// readDistrict godoc swagger
// @Summary Read one district
// @Description API to get a district by code
// @Router /districts/{code} [GET]
// @Tags District
// @Accept json
// @Produce json
// @Param code path string true "district code"
// @Success 200 {object} views.R{data=views.district}
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) readDistrict(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation, errors.New("empty branch's ID"))
		return
	}
	r, err := h.ctrl.GetDistrict(c, code)
	switch {
	case err == nil:
		h.handleSuccessResponse(c, views.District(r))
	case errors.Is(err, errs.ErrNotFound):
		h.handleSuccessResponse(c, nil)
	default:
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
	}
}

// readAllDistricts godoc swagger
// @Summary Read all district
// @Description API to get all district
// @Router /districts [GET]
// @Tags District
// @Accept json
// @Produce json
// @Success 200 {object} views.R{data=[]views.district}
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) readAllDistricts(c *gin.Context) {
	readDistrict, err := h.ctrl.ListDistrict(c)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
		return
	}
	h.handleSuccessResponse(c, views.Districts(readDistrict))
}

// updateDistrict godoc swagger
// @Summary Update district
// @Description API to update district by code
// @Router /districts/{code} [PUT]
// @Tags District
// @Accept json
// @Produce json
// @Param code path string true "district code"
// @Param request body updateDistrictParams true "update district request parameters"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) updateDistrict(c *gin.Context) {
	var req updateDistrictParams
	code := c.Param("code")

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error(err.Error())

		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation, err)
		return
	}
	model := models.District{
		CodeSOATO:    req.CodeSOATO,
		Code:         code,
		CodeGNI:      req.CodeGNI,
		CodeProvince: req.CodeProvince,
		Title: models.Languages{
			Uz: req.Title.Uz,
			Ru: req.Title.Ru,
			En: req.Title.En,
		},
		ActivatedAt: nil,
		State:       false,
	}
	if req.IsActive {
		now := time.Now()
		model.ActivatedAt = &now
	}
	err := h.ctrl.UpdateDistrict(c, model)
	switch {
	case err == nil:
		h.handleSuccessResponse(c, views.District(model))
	case errors.Is(err, errs.ErrNotFound):
		h.handleErrorResponse(c, http.StatusNotFound, status.ErrorCodeDB, err)
	default:
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
	}
}

// deleteDistrict godoc swagger
// @Summary Delete district
// @Description API to deactivate district by code
// @Router /districts/{code} [DELETE]
// @Tags District
// @Accept json
// @Produce json
// @Param code path string true "district code"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) deleteDistrict(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation,
			errors.New("empty branch ID"))
		return
	}
	err := h.ctrl.DeleteDistrict(c, code)
	switch {
	case err == nil:
		h.handleSuccessResponse(c, nil)
	case errors.Is(err, errs.ErrNotFound):
		h.handleErrorResponse(c, http.StatusNotFound, status.ErrorCodeDB, err)
	default:
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
	}
}
