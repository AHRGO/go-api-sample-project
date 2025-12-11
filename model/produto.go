package model

type Produto struct {
	ID int `json:"id_produto"`
	Name string `json:"nome"`
	Price float64 `json:"preco"`
}