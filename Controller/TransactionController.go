package Controller

import (
	model "go-test/Model"
	security "go-test/Security"
	services "go-test/Services"
	"net/http"

	"github.com/gin-gonic/gin"
)
type TransactionController struct {
	transService services.ITransactionService
}

func NewTransactionController(transService services.ITransactionService) *TransactionController{
	return &TransactionController{transService: transService}
}
func (a *TransactionController) CreateNewTransaction(c *gin.Context) {

	var transaction model.Transaction
	err := c.BindJSON(&transaction)
    if err != nil {
        c.AbortWithStatus(http.StatusBadRequest)
        return
    }

    user_id, err := security.ExtractTokenID(c)
	transaction.CustomerID=user_id
  

    res, err := a.transService.CreateNewTransaction(c.Request.Context(),&transaction)
    if err != nil {
        c.AbortWithStatus(http.StatusInternalServerError)
        return
    }

    c.JSON(http.StatusOK, res )
}
func (a *TransactionController) GetAllTransaction(c *gin.Context) {

    res, err := a.transService.GetAllTransaction(c);
    if err != nil {
        c.AbortWithStatus(http.StatusInternalServerError)
        return
    }

    c.JSON(http.StatusOK, res )
}

