package main

import "fmt"

func main() {
  x := 42
  y := 8

  x += y
  fmt.Println(x)
  y -= x
  fmt.Println(y)
  x /=y
  fmt.Println(x)
      
  x *= y
  fmt.Println(x)
}

