package main

import (
	"api-sample/controller"
	"api-sample/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	ProdutoUseCase := usecase.NewProdutoUseCase()
	ProdutoController := controller.NewProdutoController(ProdutoUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(
			200,
			gin.H {
				"message": "pong",
			})
	})

	server.GET("/produtos", ProdutoController.GetProdutos)

	server.Run(":8019")
}