package main

    import (
      "fmt"
      "conv"
    )



    func main(){
    fmt.Printf("Brrrr! %v\n", conv.AbsoluteZeroC) // "Brrrr! -273.15°C"
    fmt.Println(conv.CToF(conv.BoilingC)) // "212°F"

    }
