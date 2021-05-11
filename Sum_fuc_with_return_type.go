package main

import "fmt"

func sum(a,b,c int) int{
    return a+b+c
}

func main() {
  result := sum(42, 8,67)
  fmt.Println(result)
}
