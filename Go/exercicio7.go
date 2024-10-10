package main

import "fmt"

type Professor struct {
    Nome string
}

type Escola struct {
    Professores []Professor
}

func (e *Escola) AdicionarProfessor(p Professor) {
    e.Professores = append(e.Professores, p)
}

func main() {
    escola := Escola{}
    escola.AdicionarProfessor(Professor{"Maria"})
    escola.AdicionarProfessor(Professor{"José"})

    for _, professor := range escola.Professores {
        fmt.Println(professor.Nome)
    }
}
