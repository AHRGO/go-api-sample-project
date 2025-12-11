package main

import (
	"api-sample/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	ProdutoController := controller.NewProdutoController()

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