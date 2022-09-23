package rest

import (
	"apelsin/errs"
	"apelsin/models"
	"apelsin/pkg/status"
	"apelsin/rest/views"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type offerParams struct {
	Code           string    `json:"code" binding:"required"`
	Name           Languages `json:"name" binding:"required"`
	ShouldBeSigned bool      `json:"should_Be_Signed"`
	SignatureType  string    `json:"signature_type"`
}

type updateOfferParams struct {
	Name           Languages `json:"name" binding:"required"`
	ShouldBeSigned bool      `json:"should_Be_Signed"`
	SignatureType  string    `json:"signature_type"`
}

// createOffers godoc swagger
// @Summary Create offer
// @Description API to create a offer
// @Router /offers [POST]
// @Tags Offer
// @Accept json
// @Produce json
// @Param request body offerParams true "create offer request parameters"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) createOffers(c *gin.Context) {
	var req offerParams

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error(err.Error())
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation, err)
		return
	}
	model := models.Offer{
		Code: req.Code,
		Name: models.Languages{
			Uz: req.Name.Uz,
			Ru: req.Name.Ru,
			En: req.Name.En,
		},
		ShouldBeSigned: req.ShouldBeSigned,
		SignatureType:  req.SignatureType,
	}
	if model.ShouldBeSigned && model.SignatureType == "" {
		h.log.Error("signatureType shouldn't be null")
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation, errors.New("signatureType shouldn't be null"))
		return
	}
	if err := h.ctrl.CreateOffer(c, model); err != nil {
		h.log.Error(err.Error())
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
		return
	}
	h.handleSuccessResponse(c, nil)
}

// readAllOffers godoc swagger
// @Summary Read all offer
// @Description API to get all offers
// @Router /offers [GET]
// @Tags Offer
// @Accept json
// @Produce json
// @Success 200 {object} views.R{data=[]views.offer}
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) readAllOffers(c *gin.Context) {
	readOffer, err := h.ctrl.ListOffers(c)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
		return
	}
	h.handleSuccessResponse(c, views.Offers(readOffer))
}

// readOffer godoc swagger
// @Summary Read one offer
// @Description API to get a offer by code
// @Router /offers/{code} [GET]
// @Tags Offer
// @Accept json
// @Produce json
// @Param code path string true "offer code"
// @Success 200 {object} views.R{data=views.offer}
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) readOffer(c *gin.Context) {
	code := c.Param("code")

	if code == "" {
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation, errors.New("empty branch's ID"))
		return
	}
	r, err := h.ctrl.GetOffer(c, code)
	switch {
	case err == nil:
		h.handleSuccessResponse(c, views.Offer(r))
	case errors.Is(err, errs.ErrNotFound):
		h.handleSuccessResponse(c, nil)
	default:
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
	}
}

// updateOffer godoc swagger
// @Summary Update offer
// @Description API to update offer by code
// @Router /offers/{code} [PUT]
// @Tags Offer
// @Accept json
// @Produce json
// @Param code path string true "offer code"
// @Param request body updateOfferParams true "update offer request parameters"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) updateOffer(c *gin.Context) {
	var req updateOfferParams
	code := c.Param("code")

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error(err.Error())

		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation, err)
		return
	}
	model := models.Offer{
		Code: code,
		Name: models.Languages{
			Uz: req.Name.Uz,
			Ru: req.Name.Ru,
			En: req.Name.En,
		},
		ShouldBeSigned: req.ShouldBeSigned,
		SignatureType:  req.SignatureType,
	}
	if model.ShouldBeSigned && model.SignatureType == "" {
		h.log.Error("signatureType shouldn't be null")
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation, errors.New("signatureType shouldn't be null"))
		return
	}
	err := h.ctrl.UpdateOffer(c, model)
	switch {
	case err == nil:
		h.handleSuccessResponse(c, views.Offer(model))
	case errors.Is(err, errs.ErrNotFound):
		h.handleErrorResponse(c, http.StatusNotFound, status.ErrorCodeDB, err)
	default:
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
	}
}

// deleteOffer godoc swagger
// @Summary Delete offer
// @Description API to deactivate offer by code
// @Router /offers/{code} [DELETE]
// @Tags Offer
// @Accept json
// @Produce json
// @Param code path string true "offer code"
// @Success 200 {object} views.R
// @Failure 422 {object} views.R
// @Failure 500 {object} views.R
func (h *handler) deleteOffer(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		h.handleErrorResponse(c, http.StatusUnprocessableEntity, status.ErrorCodeValidation, errors.New("empty branch ID"))
		return
	}
	err := h.ctrl.DeleteOffer(c, code)
	switch {
	case err == nil:
		h.handleSuccessResponse(c, nil)
	case errors.Is(err, errs.ErrNotFound):
		h.handleErrorResponse(c, http.StatusNotFound, status.ErrorCodeDB, err)
	default:
		h.handleErrorResponse(c, http.StatusInternalServerError, status.ErrorCodeDB, err)
	}
}
