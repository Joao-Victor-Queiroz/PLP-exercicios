package main

import "sync"

type Configuracao struct{}

var instance *Configuracao
var once sync.Once

func GetConfiguracao() *Configuracao {
    once.Do(func() {
        instance = &Configuracao{}
    })
    return instance
}

func main() {
    config1 := GetConfiguracao()
    config2 := GetConfiguracao()
    fmt.Println(config1 == config2)  // true
}
