//type declarations
package main

import(
  "fmt"
  )

  type Celsius float64
  type Fahrenheit float64

  const(
    AbsoulteZeroC Celsius = -273.15
    FreezingC  Celsius = 0
    BoilingC   Celsius = 100
  )

  func CToF(c Celsius) Fahrenheit {return Fahrenheit(c*9/5 + 32) }
  func FToC(f Fahrenheit) Celsius {return Celsius((f-32)*5/9)}
  func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
  func main(){

    // fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
    // boilingF := CToF(BoilingC)
    // fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F
    // fmt.Printf("%g\n", boilingF-FreezingC) // este caso da erro
    var c Celsius
    // var f Fahrenheit

    //fmt.Println(c == 0)
    //fmt.Println(f >= 0)
    //fmt.Println(c == f) // vai dar erro, pois são tipos diferentes
    //fmt.Println(c == Celsius(f))

    c = FToC(212.0)
    fmt.Println(c.String())
    fmt.Printf("%v\n", c)
    fmt.Printf("%s\n", c)
    fmt.Println(c)
    fmt.Printf("%g\n", c)
    fmt.Println(float64(c))
}
