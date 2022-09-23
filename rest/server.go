package rest

import (
	"apelsin/config"
	"apelsin/pkg/logger"
	"apelsin/storage"

	"github.com/gin-gonic/gin"
)

// NewAPI ...
func NewAPI(cfg config.Config, log logger.Logger, r *gin.Engine, store storage.Storage) *gin.Engine {
	h := newHandler(cfg, log, store)
	endpoints(r, h)

	return r
}
