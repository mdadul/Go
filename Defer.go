package main

import "fmt"

func welcome() {
  fmt.Println("Welcome")
}

func main() {
  defer fmt.Print("hello")
  defer welcome()
  fmt.Println("Hey")
}
