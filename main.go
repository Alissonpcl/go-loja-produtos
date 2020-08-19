package main

import (
	"github.com/alissonpcl/go-loja-produtos/routes"
	"net/http"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
