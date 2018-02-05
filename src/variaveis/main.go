package main

import (
	"fmt"
)

//Struct tipo C

type Pessoas struct {
	Nome          string
	Idade         int
	PossuiEmprego bool
}

//METODO PEGANDO STRUCT COMO PARAMETRO
func (p Pessoas) retornaInfo() string {

	return fmt.Sprintf("Nome: %s Idade: %d", p.Nome, p.Idade)

}

//trabalhando com ponteiros

func ponteiro(s *string) {
	*s = "S PELO PONTEIRO"

}

func main() {

	//variavel normal
	var teste string
	teste = "dei valor"
	fmt.Println("PRINT VARIAVEIS NORMAIS")
	fmt.Println(teste)
	fmt.Println("saasas")
	fmt.Println("")

	//array
	testeArray := [...]int{1, 2, 3, 4, 5}
	fmt.Println("PRINT Arrays")
	fmt.Println(testeArray)
	fmt.Println(testeArray[2])
	fmt.Println("")

	//mapa
	cidades := map[string]string{"nome": "barueri", "idade": "547 sei la"}
	fmt.Println("PRINT Mapa")
	fmt.Println(cidades["nome"])
	fmt.Println(cidades)
	fmt.Println("")

	//slice ou lista
	lista := []string{"valor 1", "1", "abc"}
	//var lista [] string
	fmt.Println("PRINT Lista")
	fmt.Println(lista)
	fmt.Println("")

	// chamando estutura

	Bruno := Pessoas{Nome: "Bruno Campos Barone", Idade: 90}
	fmt.Println("PRINT Elementos da struct")
	fmt.Println(Bruno.Nome)
	fmt.Println("")

	//usando metodo que chama strutura

	fmt.Println("Usando metodo que chama estrutura")
	fmt.Println(Bruno.retornaInfo())
	fmt.Println("")

	//usando retorno por ponteiro

	var s = "MAIN"
	fmt.Println("valor do S declarado no main " + s)
	ponteiro(&s)

	fmt.Println(s)
	fmt.Println("")
}
