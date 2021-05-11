package main

import "fmt"

type student struct {
  name string
  age int
}

func main() {
  x := student{"Emdadul", 20}
  p := &x

  fmt.Println(p.age)
}
