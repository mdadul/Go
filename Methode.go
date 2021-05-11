package main

import "fmt"

type Student struct{
    name string
    id string 
    age int
}

func (x Student) show() {
    fmt.Println(x.name)
    fmt.Println(x.id)
    fmt.Println(x.age)
}

func main() {
  X := Student{"Emdadul","C201041",20}
  Y := Student{"Minhaz","C211091",20}
  X.show()
  Y.show()
}
