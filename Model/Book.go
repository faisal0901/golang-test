package Model

import "time"

type Book struct {
	Id          uint
	Title       string
	Description string
	Price       int
	CreatedAt   time.Time
	UpdatedAt   time.Time 
}