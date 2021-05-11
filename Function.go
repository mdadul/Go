package main

import "fmt"
func add(a int, b int){
    x:=a+b
    fmt.Println(x)
}
func main() {
  var a,b int
  fmt.Scanln(a)
  fmt.Scanln(b)
  
  add(a,b)
  add(5,78)
}
