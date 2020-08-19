package produtos

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos, err := BuscaTodosOsProdutos()
	if err != nil {
		panic(err.Error())
	}
	tmpl.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println(err)
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println(err)
		}

		CriaNovoProduto(nome, descricao, precoConvertido, quantidadeConvertida)
		http.Redirect(w, r, "/", 301)

	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	productIdInt, _ := strconv.Atoi(productId)
	DeleteProduct(productIdInt)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	productIdInt, _ := strconv.Atoi(productId)
	product := BuscarProduto(productIdInt)
	tmpl.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertido, err := strconv.Atoi(id)
		if err != nil {
			log.Println(err)
		}

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println(err)
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println(err)
		}

		AtualizaProduto(idConvertido, nome, descricao, precoConvertido, quantidadeConvertida)
		http.Redirect(w, r, "/", 301)
	}
}
