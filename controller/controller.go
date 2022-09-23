package controller

import (
	"apelsin/pkg/logger"
	"apelsin/storage"
)

type Controller struct {
	store storage.Storage
	logger.Logger
}

func New(store storage.Storage, logger logger.Logger) *Controller {
	return &Controller{store: store, Logger: logger}
}
