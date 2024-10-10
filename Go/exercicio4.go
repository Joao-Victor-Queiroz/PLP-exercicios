package main

import "fmt"

type Animal interface {
    Som() string
}

type Cachorro struct{}
func (Cachorro) Som() string {
    return "Au Au"
}

type Gato struct{}
func (Gato) Som() string {
    return "Miau"
}

func main() {
    var animais []Animal = []Animal{Cachorro{}, Gato{}}
    for _, animal := range animais {
        fmt.Println(animal.Som())
    }
}
