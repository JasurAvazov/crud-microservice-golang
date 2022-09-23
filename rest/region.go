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

type regionParams struct {
	Code     string    `json:"code" binding:"required"`
	Title    Languages `json:"title" binding:"required"`
	IsActive bool      `json:"is_active" binding:"required"`
	Rank     string    `json:"rank" binding:"required"`
}
type updateRegionParams struct {
	Title    Languages `json:"title" binding:"required"`
	IsActive bool      `json:"is_active" binding:"required"`
	Rank     string    `json:"rank" binding:"required"`
}

// createRegion godoc swagger
// @Summary Create region
// @Description API to create a region
// @Router /regions [POST]
// @Tags Region
// @Accept json
// @Produce json
// @Param request body regionParams true "create region request parameters"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) createRegion(c *gin.Context) {
	var req regionParams

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error(err.Error())

		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation, err)
		return
	}
	model := models.Region{
		Code: req.Code,
		Title: models.Languages{
			Uz: req.Title.Uz,
			Ru: req.Title.Ru,
			En: req.Title.En,
		},
		Rank:          req.Rank,
		ActivatedAt:   nil,
		DeactivatedAt: nil,
		State:         false,
	}
	if req.IsActive {
		now := time.Now()
		model.ActivatedAt = &now
		model.DeactivatedAt = nil
	}
	if err := h.ctrl.CreateRegion(c, model); err != nil {
		h.log.Error(err.Error())
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
		return
	}

	h.handleSuccessResponse(c, nil)
}

// readRegion godoc swagger
// @Summary Read one region
// @Description API to get a region by code
// @Router /regions/{code} [GET]
// @Tags Region
// @Accept json
// @Produce json
// @Param code path string true "region code"
// @Success 200 {object} views.R{data=views.region}
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) readRegion(c *gin.Context) {
	code := c.Param("code")

	if code == "" {
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation, errors.New("empty branch's ID"))
		return
	}
	r, err := h.ctrl.GetRegion(c, code)
	switch {
	case err == nil:
		h.handleSuccessResponse(c, views.Region(r))
	case errors.Is(err, errs.ErrNotFound):
		h.handleSuccessResponse(c, nil)
	default:
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
	}
}

// readAllRegion godoc swagger
// @Summary Read all region
// @Description API to get all region
// @Router /regions [GET]
// @Tags Region
// @Accept json
// @Produce json
// @Success 200 {object} views.R{data=[]views.region}
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) readAllRegion(c *gin.Context) {
	readRegion, err := h.ctrl.ListRegions(c)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
		return
	}
	h.handleSuccessResponse(c, views.Regions(readRegion))
}

// updateRegion godoc swagger
// @Summary Update region
// @Description API to update region by code
// @Router /regions/{code} [PUT]
// @Tags Region
// @Accept json
// @Produce json
// @Param code path string true "region code"
// @Param request body updateRegionParams true "update region request parameters"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) updateRegion(c *gin.Context) {
	var req updateRegionParams
	code := c.Param("code")

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error(err.Error())

		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation, err)
		return
	}
	model := models.Region{
		Code: code,
		Title: models.Languages{
			Uz: req.Title.Uz,
			Ru: req.Title.Ru,
			En: req.Title.En,
		},
		Rank:          req.Rank,
		ActivatedAt:   nil,
		DeactivatedAt: nil,
		State:         false,
	}
	if req.IsActive {
		now := time.Now()
		model.ActivatedAt = &now
		model.DeactivatedAt = nil
	}
	err := h.ctrl.UpdateRegion(c, model)
	switch {
	case err == nil:
		h.handleSuccessResponse(c, views.Region(model))
	case errors.Is(err, errs.ErrNotFound):
		h.handleErrorResponse(c, http.StatusNotFound, status.ErrorCodeDB, err)
	default:
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
	}
}

// deleteRegion godoc swagger
// @Summary Delete region
// @Description API to deactivate region by code
// @Router /regions/{code} [DELETE]
// @Tags Region
// @Accept json
// @Produce json
// @Param code path string true "region code"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) deleteRegion(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation, errors.New("empty branch ID"))
		return
	}
	err := h.ctrl.DeleteRegion(c, code)
	switch {
	case err == nil:
		h.handleSuccessResponse(c, nil)
	case errors.Is(err, errs.ErrNotFound):
		h.handleErrorResponse(c, http.StatusNotFound, status.ErrorCodeDB, err)
	default:
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
	}
}

// importRegions godoc swagger
// @Summary Create region
// @Description API to deactivate region by code
// @Router /regions/import [PUT]
// @Tags Region
// @Accept multipart/form-data
// @Produce json
// @Param excel formData file true "Excel File in format: A:Code|B:NameRU|C:NameEn|D:NameUz|E:Rank"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) importRegions(c *gin.Context) {
	file, err := c.FormFile("excel")
	if err != nil {
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation, err)
		return
	}
	f, err := file.Open()
	if err != nil {
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation, err)
		return
	}
	regions, err := h.ctrl.ImportRegionsToDB(c, f)
	if err != nil {
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation, err)
		return
	}

	h.handleSuccessResponse(c, views.Regions(regions))
}
