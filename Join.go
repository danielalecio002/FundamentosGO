package main

import (
	"fmt"
	"os"

)

func main() {
	
	fmt.Println(strings.Join(os.Args[1:]," "))
	//caso não queira ligar para formatação
	fmt.Println(os.Args[1:])
}