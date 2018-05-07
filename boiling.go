// O seguinte programa mostra o ponto de ebulição da agua em fF e  C
package main

import  "fmt"

const ebulicaoF = 212.0

func main() {
    var f = ebulicaoF
    var c = (f-  32) * 5 / 9
    fmt.Printf("Ponto de ebulicao = %gºF or %gºC\n", f, c)

}