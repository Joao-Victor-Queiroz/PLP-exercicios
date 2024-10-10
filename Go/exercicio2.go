package main

import "fmt"

type Carro struct {
    Marca     string
    Modelo    string
    Ano       int
    Velocidade int
}

func (c *Carro) Acelerar(incremento int) {
    c.Velocidade += incremento
}

func (c *Carro) Frear(decremento int) {
    c.Velocidade -= decremento
}

func (c Carro) ExibirVelocidade() {
    fmt.Printf("Velocidade atual: %d km/h\n", c.Velocidade)
}

func main() {
    carro := Carro{"Toyota", "Corolla", 2020, 0}
    carro.Acelerar(50)
    carro.ExibirVelocidade()
    carro.Frear(20)
    carro.ExibirVelocidade()
}
