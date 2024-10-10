package main

import "fmt"

type Motor struct {
    Potencia int
}

type Carro struct {
    Marca  string
    Modelo string
    Motor  Motor
}

func main() {
    motor1 := Motor{Potencia: 150}
    carro := Carro{Marca: "Toyota", Modelo: "Corolla", Motor: motor1}
    fmt.Printf("%s %s com motor de %d HP\n", carro.Marca, carro.Modelo, carro.Motor.Potencia)
}
