package Controller

import (
	model "go-test/Model"
	security "go-test/Security"
	services "go-test/Services"
	"net/http"

	"github.com/gin-gonic/gin"
)


type AuthController struct {
	authService services.IAuthService
}


func NewAuthController(authService services.IAuthService) *AuthController {
	return &AuthController{authService: authService}
}



func (a *AuthController) Register(c *gin.Context) {
    var customer model.Customer
    err := c.BindJSON(&customer)
    if err != nil {
        c.AbortWithStatus(http.StatusBadRequest)
        return
    }

    res, err := a.authService.CreateCustomer(c.Request.Context(), &customer)
    if err != nil {
        c.AbortWithStatus(http.StatusInternalServerError)
        return
    }

 
   
    c.JSON(http.StatusOK, res )
}
func (a *AuthController) Login(c *gin.Context) {
    var customer model.Customer
    err := c.BindJSON(&customer)
    if err != nil {
        c.AbortWithStatus(http.StatusBadRequest)
        return
    }

    res, err := a.authService.LoginCustomer(c.Request.Context(), &customer)
    if err != nil {
        c.AbortWithStatus(http.StatusInternalServerError)
        return
    }

 
    c.JSON(http.StatusOK, res )
}
func (a *AuthController) Logout(c *gin.Context) {
  
 
    token:=security.ExtractToken(c)
    err:=a.authService.LogoutCustomer(c.Request.Context(),token)
    if err != nil {
        c.AbortWithStatus(http.StatusInternalServerError)
        return
    }
    c.JSON(http.StatusAccepted,err)
 
}



