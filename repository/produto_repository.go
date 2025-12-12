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

func (pr *ProdutoRepository) CreateProduto(produto model.Produto) (int, error) {

	var id int
	query, err := pr.connection.Prepare( "" +
		"INSERT INTO produto" +
		"(nome_produto, preco)" + 
		"VALUES ($1, $2) RETURNING id_produto")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}


	err = query.QueryRow(produto.Name, produto.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}
