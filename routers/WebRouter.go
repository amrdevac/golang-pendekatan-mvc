package WebRouter

import (
	handler "pustaka-api/Handler"
	BooksController "pustaka-api/controllers"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
)

func RouterApplication() {
	routes := gin.Default()

	v1 := routes.Group("/v1")

	v1.GET("/", handler.HelloHandler)
	v1.GET("/:id", handler.RootHandler)
	v1.GET("/query-handler", handler.QueryHandler)
	v1.POST("/post-input", BooksController.PostInput)

	routes.Run("127.0.0.1:8080")

}
