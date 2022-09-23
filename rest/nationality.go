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

type nationalityParams struct {
	Code  string    `json:"code" binding:"required"`
	Title Languages `json:"title" binding:"required"`
	State bool      `json:"state" binding:"required"`
}
type updateNationalityParams struct {
	Title Languages `json:"title" binding:"required"`
	State bool      `json:"state" binding:"required"`
}

// createNationality godoc swagger
// @Summary Create nationality
// @Description API to create a nationality
// @Router /nationalities [POST]
// @Tags Nationality
// @Accept json
// @Produce json
// @Param request body nationalityParams true "create nationality request parameters"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) createNationality(c *gin.Context) {
	var req nationalityParams

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error(err.Error())

		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation, err)
		return
	}
	model := models.Nationality{
		Code: req.Code,
		Title: models.Languages{
			Uz: req.Title.Uz,
			Ru: req.Title.Ru,
			En: req.Title.En,
		},
		ActivatedAt:   nil,
		DeactivatedAt: nil,
		State:         false,
	}
	if req.State {
		now := time.Now()
		model.ActivatedAt = &now
		model.DeactivatedAt = nil
	}
	if err := h.ctrl.CreateNationality(c, model); err != nil {
		h.log.Error(err.Error())
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
		return
	}

	h.handleSuccessResponse(c, nil)
}

// readNationality godoc swagger
// @Summary Read one nationality
// @Description API to get a nationality by code
// @Router /nationalities/{code} [GET]
// @Tags Nationality
// @Accept json
// @Produce json
// @Param code path string true "nationality code"
// @Success 200 {object} views.R{data=views.nationality}
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) readNationality(c *gin.Context) {
	code := c.Param("code")

	if code == "" {
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation, errors.New("empty branch's ID"))
		return
	}
	r, err := h.ctrl.GetNationality(c, code)
	switch {
	case err == nil:
		h.handleSuccessResponse(c, views.Nationality(r))
	case errors.Is(err, errs.ErrNotFound):
		h.handleSuccessResponse(c, nil)
	default:
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
	}
}

// readAllNationalities godoc swagger
// @Summary Read all nationality
// @Description API to get all nationalities
// @Router /nationalities [GET]
// @Tags Nationality
// @Accept json
// @Produce json
// @Success 200 {object} views.R{data=[]views.nationality}
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) readAllNationalities(c *gin.Context) {
	readNationality, err := h.ctrl.ListNationalities(c)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
		return
	}
	h.handleSuccessResponse(c, views.Nationalities(readNationality))
}

// updateNationality godoc swagger
// @Summary Update nationality
// @Description API to update nationality by code
// @Router /nationalities/{code} [PUT]
// @Tags Nationality
// @Accept json
// @Produce json
// @Param code path string true "nationality code"
// @Param request body updateNationalityParams true "update nationality request parameters"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) updateNationality(c *gin.Context) {
	var req updateNationalityParams
	code := c.Param("code")

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error(err.Error())
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation, err)
		return
	}
	model := models.Nationality{
		Code: code,
		Title: models.Languages{
			Uz: req.Title.Uz,
			Ru: req.Title.Ru,
			En: req.Title.En,
		},
		ActivatedAt:   nil,
		DeactivatedAt: nil,
		State:         false,
	}
	if req.State {
		now := time.Now()
		model.ActivatedAt = &now
		model.DeactivatedAt = nil
	}
	err := h.ctrl.UpdateNationality(c, model)
	switch {
	case err == nil:
		h.handleSuccessResponse(c, views.Nationality(model))
	case errors.Is(err, errs.ErrNotFound):
		h.handleErrorResponse(c, http.StatusNotFound, status.ErrorCodeDB, err)
	default:
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
	}
}

// deleteNationality godoc swagger
// @Summary Delete nationality
// @Description API to deactivate nationality by code
// @Router /nationalities/{code} [DELETE]
// @Tags Nationality
// @Accept json
// @Produce json
// @Param code path string true "nationality code"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) deleteNationality(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation,
			errors.New("empty branch ID"))
		return
	}
	err := h.ctrl.DeleteNationality(c, code)
	switch {
	case err == nil:
		h.handleSuccessResponse(c, nil)
	case errors.Is(err, errs.ErrNotFound):
		h.handleErrorResponse(c, http.StatusNotFound, status.ErrorCodeDB, err)
	default:
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
	}
}

// importNationalities godoc swagger
// @Summary Create nationality
// @Description API to deactivate nationality by code
// @Router /nationalities/import [PUT]
// @Tags Nationality
// @Accept multipart/form-data
// @Produce json
// @Param excel formData file true "Excel File in format: A:Code|B:NameRU|C:NameEn|D:NameUz"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) importNationalities(c *gin.Context) {
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
	nationalities, err := h.ctrl.ImportNationalitiesToDB(c, f)
	if err != nil {
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation, err)
		return
	}
	h.handleSuccessResponse(c, views.Nationalities(nationalities))
}
