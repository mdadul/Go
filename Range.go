package main

import "fmt"

func main() {
  a := make([]int, 50)
  a[1] = 2
  a[2] = 3
  a[0] = 78

  for i, v:= range a{
      fmt.Println(i+1,v)
  }
}
