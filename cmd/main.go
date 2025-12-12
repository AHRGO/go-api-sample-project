package main

import (
	"api-sample/controller"
	"api-sample/db"
	"api-sample/repository"
	"api-sample/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConn, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	ProdutoRepository := repository.NewProdutoRepository(dbConn)
	ProdutoUseCase := usecase.NewProdutoUseCase(ProdutoRepository)
	ProdutoController := controller.NewProdutoController(ProdutoUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(
			200,
			gin.H {
				"message": "pong",
			})
	})

	server.GET("/produtos", ProdutoController.GetProdutos)

	server.POST("/produto", ProdutoController.CreateProduto)

	server.Run(":8019")
}