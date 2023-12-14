package main

import (
	"fmt"

	"github.com/mylenaimai/go-api/database"
	"github.com/mylenaimai/go-api/models"
	"github.com/mylenaimai/go-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "nome1", Historia: "hist1"},
		{Id: 2, Nome: "nome2", Historia: "hist2"},
	}
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor em go")
	routes.HandleRequest()
}
