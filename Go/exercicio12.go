package main

import "fmt"

type Produto struct {
    Preco float64
}

func SomarProdutos(p1, p2 Produto) Produto {
    return Produto{Preco: p1.Preco + p2.Preco}
}

func main() {
    produto1 := Produto{Preco: 50}
    produto2 := Produto{Preco: 30}
    produto3 := SomarProdutos(produto1, produto2)
    fmt.Println(produto3.Preco)
}
