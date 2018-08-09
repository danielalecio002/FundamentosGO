package main

import "fmt"

func duplicar(strings []string) [] string{
var rs []string

for i := 1; i < len(strings) ; i++ {
               if strings[i] != strings[i - 1]{
                   rs = append(rs,strings[i - 1])
              }
              if len(strings) - 1 == i{
                  rs = append(rs,strings[i])
              }
}
return rs
}
func main(){
data := []string {"d","jkhgf","jkhgf","maria","dan","dan"} //"d","jkhgf","maria","dan"
fmt.Printf("%q", duplicar(data))
}