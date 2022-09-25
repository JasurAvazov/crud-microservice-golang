package rest

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func endpoints(r *gin.Engine, h *handler) {
	url := ginSwagger.URL("../swagger/doc.json")

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.POST("/districts", h.createDistrict)
	r.GET("/districts", h.readAllDistricts)
	r.GET("/districts/:code", h.readDistrict)
	r.PUT("/districts/:code", h.updateDistrict)
	r.DELETE("/districts/:code", h.deleteDistrict)

}
