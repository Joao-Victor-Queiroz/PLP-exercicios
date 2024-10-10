package main

import "fmt"

func Somar(a, b int) int {
    return a + b
}

func SomarTres(a, b, c int) int {
    return a + b + c
}

func main() {
    fmt.Println(Somar(1, 2))
    fmt.Println(SomarTres(1, 2, 3))
}
