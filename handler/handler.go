package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name": "fahmi amrullah",
	})
}

func RootHandler(ctx *gin.Context) {
	name := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"name": "tia andrian",
		"id":   name,
	})
}

func QueryHandler(ctx *gin.Context) {
	title := ctx.Query("title")

	ctx.JSON(http.StatusOK, gin.H{
		"request": title,
	})
}
