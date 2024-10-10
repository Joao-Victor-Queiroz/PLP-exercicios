package main

import "fmt"

type Funcionario interface {
    CalcularSalario() float64
}

type FuncionarioHorista struct {
    HorasTrabalhadas int
    ValorPorHora     float64
}

func (f FuncionarioHorista) CalcularSalario() float64 {
    return float64(f.HorasTrabalhadas) * f.ValorPorHora
}

type FuncionarioAssalariado struct {
    Salario float64
}

func (f FuncionarioAssalariado) CalcularSalario() float64 {
    return f.Salario
}

func main() {
    funcionario1 := FuncionarioHorista{160, 25}
    funcionario2 := FuncionarioAssalariado{3000}

    fmt.Println(funcionario1.CalcularSalario())
    fmt.Println(funcionario2.CalcularSalario())
}
