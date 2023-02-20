package Controller

import (
	services "go-test/Services"
	"net/http"

	"github.com/gin-gonic/gin"
)
type BookController struct {
	bookService services.IBookService
}

func NewBookController(bookService services.IBookService) *BookController {
	return &BookController{bookService: bookService}
}

func (u *BookController)GetAllBooks(c *gin.Context)  {
	books, err := u.bookService.GetAllBooks(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, books)
}