package router

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/recipes", recipesHandler.ListRecipesHandler)

}
