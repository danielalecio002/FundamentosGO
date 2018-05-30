package main

import (
    "fmt"
)

func comma(s string ) string {
    n := len(s)
    if n <= 3{
        return s
    }
    return comma(s[:n-3]) + "," + s[n-3:]
}

func main() {
    s1 := "12345"
   // s2 := "abc"

    fmt.Printf("%s",comma(s1));

}