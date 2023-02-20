package dtos

type BookInput struct {
	Title string `json:"title" binding:"required"`
	Price int    `json:"price" binding:"required"`
}