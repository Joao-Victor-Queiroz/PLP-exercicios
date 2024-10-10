package main

import "fmt"

type Carro struct {
    Marca  string
    Modelo string
    Ano    int
}

func main() {
    carro1 := Carro{"Toyota", "Corolla", 2020}
    carro2 := Carro{"Honda", "Civic", 2019}
    carro3 := Carro{"Ford", "Focus", 2021}

    for _, carro := range []Carro{carro1, carro2, carro3} {
        fmt.Printf("%s %s %d\n", carro.Marca, carro.Modelo, carro.Ano)
    }
}
