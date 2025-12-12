package repository

import (
	"api-sample/model"
	"database/sql"
	"fmt"
)

type ProdutoRepository struct {
	connection *sql.DB
}

func NewProdutoRepository(connection *sql.DB) ProdutoRepository {
	return ProdutoRepository {
		connection: connection,
	}
}

func (pr *ProdutoRepository) GetAllProdutos() ([]model.Produto, error) {
	query := "SELECT id_produto, nome_produto, preco FROM produto"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Produto{}, err
	}

	var produtoList []model.Produto
	var produtoObj model.Produto

	for rows.Next() {
		err = rows.Scan(
			&produtoObj.ID,
			&produtoObj.Name,
			&produtoObj.Price)
		
		if err != nil {
			fmt.Println(err)
			return []model.Produto{}, err
		}

		produtoList = append(produtoList, produtoObj)
	}

	rows.Close()

	return produtoList, nil
}