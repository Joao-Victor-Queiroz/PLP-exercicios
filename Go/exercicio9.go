package main

import "fmt"

type Imprimivel interface {
    Imprimir()
}

type Relatorio struct{}
func (Relatorio) Imprimir() {
    fmt.Println("Imprimindo relatório")
}

type Contrato struct{}
func (Contrato) Imprimir() {
    fmt.Println("Imprimindo contrato")
}

func main() {
    itens := []Imprimivel{Relatorio{}, Contrato{}}
    for _, item := range itens {
        item.Imprimir()
    }
}
