package Controller

import (
	model "go-test/Model"
	security "go-test/Security"
	services "go-test/Services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
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
type LoginRequest struct {
	Email    string `json:"Email" validate:"required,email"`
	Password string `json:"Password" validate:"required,min=6"`
}
func (a *AuthController) Login(c *gin.Context) {
    var loginReq LoginRequest
	if err := c.ShouldBindWith(&loginReq, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	if err := validator.New().Struct(loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer := model.Customer{
		Email:    loginReq.Email,
		Password: loginReq.Password,
	}
    response, errResp := a.authService.LoginCustomer(c.Request.Context(), &customer)
    if errResp.Code != 0 {
        c.JSON(errResp.Code, errResp)
        return
    }

    c.JSON(http.StatusOK, response )
}
func (a *AuthController) Logout(c *gin.Context) {
  
 
    token:=security.ExtractToken(c)
    user_id, err := security.ExtractTokenID(c)
    res,err:=a.authService.LogoutCustomer(c.Request.Context(),token,user_id)
    if err != nil {
        c.AbortWithStatus(http.StatusInternalServerError)
        return
    }
    c.JSON(http.StatusOK,res)
 
}



