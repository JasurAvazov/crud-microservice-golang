package rest

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func endpoints(r *gin.Engine, h *handler) {
	url := ginSwagger.URL("../swagger/doc.json")

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.PUT("/regions/import", h.importRegions)
	r.POST("/regions", h.createRegion)
	r.GET("/regions", h.readAllRegion)
	r.GET("/regions/:code", h.readRegion)
	r.PUT("/regions/:code", h.updateRegion)
	r.DELETE("/regions/:code", h.deleteRegion)

	r.PUT("/districts/import", h.importDistricts)
	r.POST("/districts", h.createDistrict)
	r.GET("/districts", h.readAllDistricts)
	r.GET("/districts/:code", h.readDistrict)
	r.PUT("/districts/:code", h.updateDistrict)
	r.DELETE("/districts/:code", h.deleteDistrict)

	r.PUT("/nationalities/import", h.importNationalities)
	r.POST("/nationalities", h.createNationality)
	r.GET("/nationalities", h.readAllNationalities)
	r.GET("/nationalities/:code", h.readNationality)
	r.PUT("/nationalities/:code", h.updateNationality)
	r.DELETE("/nationalities/:code", h.deleteNationality)

	r.POST("/offers", h.createOffers)
	r.GET("/offers", h.readAllOffers)
	r.GET("/offers/:code", h.readOffer)
	r.PUT("/offers/:code", h.updateOffer)
	r.DELETE("/offers/:code", h.deleteOffer)
}
