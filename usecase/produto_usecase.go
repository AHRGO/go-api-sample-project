package usecase

import "api-sample/model"

type ProdutoUseCase struct {
	//Repository
}

func NewProdutoUseCase() ProdutoUseCase {
	return ProdutoUseCase{}
}

func (pu *ProdutoUseCase) GetProdutos() ([]model.Produto, error) {
	return []model.Produto{}, nil
}