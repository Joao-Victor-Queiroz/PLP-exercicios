package main

import "fmt"

type Empregado struct {
    Nome   string
    Cargo  string
    Salario float64
}

type Empresa struct {
    Empregados []Empregado
}

func (e *Empresa) AdicionarEmpregado(empregado Empregado) {
    e.Empregados = append(e.Empregados, empregado)
}

func main() {
    empresa := Empresa{}
    empresa.AdicionarEmpregado(Empregado{"João", "Gerente", 5000})
    empresa.AdicionarEmpregado(Empregado{"Ana", "Desenvolvedora", 4000})

    for _, empregado := range empresa.Empregados {
        fmt.Printf("%s, %s, R$%.2f\n", empregado.Nome, empregado.Cargo, empregado.Salario)
    }
}
