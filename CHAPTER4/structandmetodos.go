package main

import "fmt"

type anotacoes struct{
titulo string
texto string
data string
}
func (A * anotacoes) escrever() string {
    return "         " + A.titulo + "\n" + A.texto + "\n" + A.data
}

func main() {
    anot := anotacoes{"Struct","Utilizandos struct e metodos em golang ","10/07/2018"}
    fmt.Println(anot.escrever())
}