// programa que converte de graus F para C
package main

import(
    "fmt"
)

func main(){
    const freezinF, boilingF = 32.0 ,212.0

    fmt.Printf("%gºF  = %gºC\n", freezinF, FtoC(freezinF))
    fmt.Printf("%gºF = %gºC\n", boilingF, FtoC(boilingF))
}
func FtoC(f float64) float64{
    return (f - 32) * 5 / 9
}