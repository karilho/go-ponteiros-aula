package main

import "fmt"

type Pessoa struct {
	Nome string
	Idade int 
}

func main() {
	Pessoa1 := Pessoa{}
	Pessoa1.Nome = "Lucas"
	Pessoa1.Idade = 20

	fmt.Println(&Pessoa1)
	fazerAniversario(&Pessoa1)
	fmt.Println(Pessoa1)

}

func fazerAniversario(pessoa *Pessoa){
	pessoa.Idade++
	fmt.Println("rodando dentro da função", pessoa.Idade)
}


