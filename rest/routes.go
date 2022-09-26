package rest

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func endpoints(r *gin.Engine, h *handler) {
	url := ginSwagger.URL("../swagger/doc.json")

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.POST("/customer", h.createCustomer)
	r.GET("/customers", h.readAllCustomers)
	r.GET("/customer/:code", h.readCustomer)
	r.PUT("/customer/:code", h.updateCustomer)
	r.DELETE("/customer/:code", h.deleteCustomer)

	r.POST("/order", h.createOrder)
	r.GET("/orders", h.readAllOrders)
	r.GET("/order/:code", h.readOrder)
	r.PUT("/order/:code", h.updateOrder)
	r.DELETE("/order/:code", h.deleteOrder)

	r.POST("/category", h.createCategory)
	r.GET("/categories", h.readAllCategories)
	r.GET("/category/:code", h.readCategory)
	r.PUT("/category/:code", h.updateCategory)
	r.DELETE("/category/:code", h.deleteCategory)

	r.POST("/product", h.createProduct)
	r.GET("/products", h.readAllProducts)
	r.GET("/product/:code", h.readProduct)
	r.PUT("/product/:code", h.updateProduct)
	r.DELETE("/product/:code", h.deleteProduct)

	r.POST("/detail", h.createDetail)
	r.GET("/details", h.readAllDetails)
	r.GET("/detail/:code", h.readDetail)
	r.PUT("/detail/:code", h.updateDetail)
	r.DELETE("/detail/:code", h.deleteDetail)

}
