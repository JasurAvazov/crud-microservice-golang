package rest

import (
	"apelsin/controller"
	"apelsin/pkg/status"
	"apelsin/rest/views"
	"apelsin/storage"

	"net/http"
	"strconv"

	"apelsin/config"
	"apelsin/pkg/logger"

	"github.com/gin-gonic/gin"
)

type handler struct {
	cfg     config.Config
	log     logger.Logger
	storage storage.Storage
	ctrl    *controller.Controller
}

func newHandler(cfg config.Config, log logger.Logger, store storage.Storage) *handler {
	return &handler{
		cfg:     cfg,
		log:     log,
		storage: store,
		ctrl:    controller.New(store, log),
	}
}

func (h *handler) handleSuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, views.R{
		Status:    status.Success,
		ErrorCode: 0,
		ErrorNote: "",
		Data:      data,
	})
}

func (h *handler) handleErrorResponse(c *gin.Context, httpCode, errCode int, err error) {
	c.JSON(httpCode, views.R{
		Status:    status.Failure,
		ErrorCode: errCode,
		ErrorNote: err.Error(),
		Data:      nil,
	})

}

func parseOffsetQueryParam(c *gin.Context) (int, error) {
	return strconv.Atoi(c.DefaultQuery("offset", "0"))
}

func parseLimitQueryParam(c *gin.Context) (int, error) {
	return strconv.Atoi(c.DefaultQuery("limit", "10"))
}

type Languages struct {
	Uz string `json:"uz"`
	Ru string `json:"ru"`
	En string `json:"en"`
}
