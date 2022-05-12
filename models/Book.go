package models

type Book struct {
	Title string `binding:"required"`
	Price string `json:"price" binding:"required,number"`
}
