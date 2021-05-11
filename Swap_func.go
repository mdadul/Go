package main

import "fmt"

func swap(a,b int) (int,int){
    return b,a
}
func main() {
  a, b := swap(42, 8)
  fmt.Println(a,b)
  fmt.Println(b)
}
