package main

import "fmt"

type SaldoInsuficienteError struct {
    Mensagem string
}

func (e *SaldoInsuficienteError) Error() string {
    return e.Mensagem
}

type ContaBancaria struct {
    saldo float64
}

func (c *ContaBancaria) Depositar(valor float64) {
    c.saldo += valor
}

func (c *ContaBancaria) Sacar(valor float64) error {
    if valor > c.saldo {
        return &SaldoInsuficienteError{"Saldo insuficiente"}
    }
    c.saldo -= valor
    return nil
}

func main() {
    conta := ContaBancaria{}
    conta.Depositar(100)
    err := conta.Sacar(150)
    if err != nil {
        fmt.Println(err)
    }
}
