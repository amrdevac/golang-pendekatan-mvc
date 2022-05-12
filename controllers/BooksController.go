package BooksController

import (
	"fmt"
	"net/http"
	"pustaka-api/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func PostInput(ctx *gin.Context) {
	var input models.Book

	err := ctx.ShouldBindJSON(&input)

	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on filed %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"input": input.Title,
		"price": input.Price,
	})
}
