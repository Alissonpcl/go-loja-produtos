package routes

import (
	"github.com/alissonpcl/go-loja-produtos/produtos"
	"net/http"
)

func CarregaRotas()  {
	http.HandleFunc("/", produtos.Index)
	http.HandleFunc("/new", produtos.New)
	http.HandleFunc("/insert", produtos.Insert)
	http.HandleFunc("/delete", produtos.Delete)
	http.HandleFunc("/edit", produtos.Edit)
	http.HandleFunc("/update", produtos.Update)
}
