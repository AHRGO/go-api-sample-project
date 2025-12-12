package usecase

import (
	"api-sample/model"
	"api-sample/repository"
)

type ProdutoUseCase struct {
	repository repository.ProdutoRepository
}

func NewProdutoUseCase(repo repository.ProdutoRepository) ProdutoUseCase {
	return ProdutoUseCase{
		repository: repo,
	}
}

func (pu *ProdutoUseCase) GetProdutos() ([]model.Produto, error) {
	return pu.repository.GetAllProdutos()
}

func (pu *ProdutoUseCase) CreateProduto(produto model.Produto) (model.Produto, error) {
	idProduto, err := pu.repository.CreateProduto(produto)
	if err != nil {
		return model.Produto{}, err
	}

	produto.ID = idProduto

	return produto, nil
} 