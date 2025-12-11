package controller

import (
	"api-sample/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type produtoController struct {

}

func NewProdutoController() produtoController {
	return produtoController{}
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

