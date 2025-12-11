package controller

import (
	"api-sample/model"
	"api-sample/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type produtoController struct {
	produtoUseCase usecase.ProdutoUseCase
}

func NewProdutoController(usecase usecase.ProdutoUseCase) produtoController {
	return produtoController{
		produtoUseCase: usecase,
	}
}

func (prd *produtoController) GetProdutos(ctx *gin.Context) {

	produtos := []model.Produto{
		{
			ID: 1,
			Name: "Macarrão",
			Price: 7,
		},
	}

	ctx.JSON(http.StatusOK, produtos)
}

