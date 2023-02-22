package Controller

import (
	model "go-test/Model"
	services "go-test/Services"
	"net/http"

	"github.com/gin-gonic/gin"
)
type ProductController struct {
	productService services.IProductService
}

func NewProductController(productService services.IProductService) *ProductController{
	return &ProductController{productService: productService}
}
func (a *ProductController) CreateNewProduct(c *gin.Context) {
    var product model.Product
    err := c.BindJSON(&product)
    if err != nil {
        c.AbortWithStatus(http.StatusBadRequest)
        return
    }

    res, err := a.productService.CreateNewProduct(c.Request.Context(), &product)
    if err != nil {
        c.AbortWithStatus(http.StatusInternalServerError)
        return
    }

    c.JSON(http.StatusOK, res )
}