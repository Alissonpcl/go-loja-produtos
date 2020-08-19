package produtos

import (
	db2 "github.com/alissonpcl/go-loja-produtos/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func CriaNovoProduto(nome string, descricao string, preco float64, quantidade int) {
	db := db2.ConectaComBancoDeDados()
	defer db.Close()

	preparedSQL, err := db.Prepare(`insert into produtos (nome, descricao, preco, quantidade) values ($1, $2, $3, $4)`)

	if err != nil {
		panic(err)
	}

	preparedSQL.Exec(nome, descricao, preco, quantidade)

}

func AtualizaProduto(id int, nome string, descricao string, preco float64, quantidade int) {
	db := db2.ConectaComBancoDeDados()
	defer db.Close()

	preparedSQL, err := db.Prepare(`update produtos set 
											nome=$2, 
											descricao=$3, 
											preco=$4, 
											quantidade=$5 
											where id = $1`)

	if err != nil {
		panic(err)
	}

	preparedSQL.Exec(id, nome, descricao, preco, quantidade)

}

func BuscaTodosOsProdutos() ([]Produto, error) {
	db := db2.ConectaComBancoDeDados()
	defer db.Close()

	selectTodosProdutos, err := db.Query(`select id, nome, descricao, preco, quantidade 
												from produtos 
												order by nome`)
	if err != nil {
		return nil, err
	}

	p := Produto{}
	var produtos []Produto
	for selectTodosProdutos.Next() {
		err = selectTodosProdutos.Scan(&p.Id, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade)
		if err != nil {
			return nil, err
		}
		produtos = append(produtos, p)
	}

	return produtos, nil
}

func BuscarProduto(id int) Produto {
	db := db2.ConectaComBancoDeDados()
	defer db.Close()
	selectPrepared := db.QueryRow("select id, nome, descricao, preco, quantidade from produtos where id = $1", id)

	p := Produto{}
	selectPrepared.Scan(&p.Id, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade)

	return p
}

func DeleteProduct(id int) {
	db := db2.ConectaComBancoDeDados()
	defer db.Close()

	deleteProduct, err := db.Prepare("Delete from produtos where id = $1")
	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)
}
