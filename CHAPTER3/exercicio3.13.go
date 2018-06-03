package main

import (
    "fmt"
)
type size float64

const(
_           = iota
KIB  size = 1 << (10 * iota)
MIB
GIB
TIB
PIZ
EIB
ZIB
YIB
)
func main(){

fmt.Println(YIB)
}