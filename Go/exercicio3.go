package main

import "fmt"

type ContaBancaria struct {
    titular string
    saldo   float64
}

func (c *ContaBancaria) Depositar(valor float64) {
    c.saldo += valor
}

func (c *ContaBancaria) Sacar(valor float64) error {
    if valor > c.saldo {
        return fmt.Errorf("saldo insuficiente")
    }
    c.saldo -= valor
    return nil
}

func (c ContaBancaria) ObterSaldo() float64 {
    return c.saldo
}

func main() {
    conta := ContaBancaria{titular: "João"}
    conta.Depositar(100)
    fmt.Println(conta.ObterSaldo())
    err := conta.Sacar(50)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(conta.ObterSaldo())
}
