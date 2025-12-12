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

func (p *produtoController) GetProdutos(ctx *gin.Context) {
	
	produtos, err := p.produtoUseCase.GetProdutos()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, produtos)
}


func (pc *produtoController) CreateProduto(ctx *gin.Context) {
	
	var produto model.Produto
	err := ctx.BindJSON(&produto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	produtoCriado, err := pc.produtoUseCase.CreateProduto(produto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, produtoCriado)
}

