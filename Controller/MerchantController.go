package Controller

import (
	model "go-test/Model"
	services "go-test/Services"
	"net/http"

	"github.com/gin-gonic/gin"
)
type MerchantController struct {
	MerchantService services.IMerchantService
}

func NewMerchantController(MerchantService services.IMerchantService) *MerchantController{
	return &MerchantController{MerchantService: MerchantService}
}
func (a *MerchantController) GetAllMerchant(c *gin.Context) {
	res, err := a.MerchantService.GetAllMerchant(c);
    if err != nil {
        c.AbortWithStatus(http.StatusInternalServerError)
        return
    }

    c.JSON(http.StatusOK, res )
}
func (a *MerchantController) CreateNewMerchant(c *gin.Context) {
	var merchant model.Merchant
	err := c.BindJSON(&merchant)
    if err != nil {
        c.AbortWithStatus(http.StatusBadRequest)
        return
    }
	res, err := a.MerchantService.CreateMerchant(c.Request.Context(),&merchant);
    if err != nil {
        c.AbortWithStatus(http.StatusInternalServerError)
        return
    }

    c.JSON(http.StatusOK, res )
}