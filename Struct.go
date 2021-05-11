package main

import "fmt"

type Student struct {
    name string
    age int 
    section string 
    id string
}
func main() {
  X := Student{"Emdadul",20,"A","C201041"}
  Y := Student{"Minhaz",20,"C","C211091"}
  
  fmt.Println(X.name)
  fmt.Println(X.age)
  fmt.Println(X.section)
  fmt.Println(X.id)
  
  fmt.Println(Y.name)
  fmt.Println(Y.age)
  fmt.Println(Y.section)
  fmt.Println(Y.id)
  
  
}
