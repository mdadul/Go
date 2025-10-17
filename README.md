+ [Go](#go)
+ [Hello world](#hello-world)
+ [Package](#package)
+ [Import](#import)
+ [Variable](#variables)
+ [Data Type](#data-type)
+ [Arithmetic Operator](#arithmetic-operator)
+ [Assignment Operators](#assignment-operators)
+ [Constants](#constants)
+ [If-Else Statements](#if-else-statements)
+ [For Loops](#for-loops)
+ [Functions](#functions)
+ [Functions with Return Values](#functions-with-return-values)
+ [Multiple Return Values](#multiple-return-values)
+ [Structs](#structs)
+ [Methods](#methods)
+ [Pointers](#pointers)
+ [Slices](#slices)
+ [Range](#range)
+ [String Operations](#string-operations)
+ [Defer Statement](#defer-statement)
+ [Make Function](#make-function)
+ [Append Function](#append-function)
+ [Time Package](#time-package)

# Go
Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.

Go was designed at Google in 2007 to improve programming productivity in an era of multicore, networked machines and large codebases.



# Hello world
```Go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
} 
```
# Package 
Every Go program is made up of packages.

A Go program starts running in the main package.

This is why we need to declare our code as the main package -- to make it run and create the output:


# Import 
Apart from main, Go has many packages that can be imported and used in the code to accomplish different tasks.

One of the most popular packages is "fmt", which stands for format, and provides input and output functionality.

To import a package, we use the import statement:


# Variables

Variables are used to store values.

In Go, the var keyword is used to declare variables.

```Go
	var x int
```
The code above declares a variable named i of type int.
int stands for integer and represents whole numbers.

We can assign the variable a value and output it:

```Go
	var x int = 8
```
You can also declare multiple variables on one line and assign values to them:

```Go
	var i,j int = 87,89
```
If you assign a value to the variable, you can omit the type declaration, as Go will automatically take the type of the value assigned:

```Go
	var i,name = 89,"Go"
```
Go supports short variable declaration using :=.

```Go
	x := 89
	name := "Emdadul"
```


## Data Type ##

Let's see what other common types Go supports.
- float32 - a single-precision floating point value.
- float64 - a double-precision floating point value.
- string - a text value.
- bool - Boolean true or false.

```Go
package main

import "fmt"

func main() {
  var x int = 76742
  var y float32 = 1.3877
  var name string = "Emdadul"
  var flag bool = true

  fmt.Println(name)
  fmt.Println(x)
  fmt.Println(y)
  fmt.Println(flag)
}
```
## Arithmetic Operator ##
Go supports all of the common arithmetic operators.
(+ - * / %) 

```Go
package main

import "fmt"

func main() {

  x := 45
  y := 7
  
  //Addition
  add := x+y
  fmt.Println(add)
  
  //Subtraction
  sub := x-y
  fmt.Println(sub)
  
  //Multipication
  mul := x*y
  fmt.Println(mul)
  
  //Division 
  div := x/y
  fmt.Println(div)
      
  // Modulus, results in the remainder of the division
  rem := x % y
  fmt.Println(rem) 
}

```

## Assignment Operators

Assignment operators are used to assign values to variables and perform operations at the same time.

Common assignment operators in Go:
- `+=` - Addition assignment
- `-=` - Subtraction assignment  
- `*=` - Multiplication assignment
- `/=` - Division assignment
- `%=` - Modulus assignment

```Go
package main

import "fmt"

func main() {
  x := 42
  y := 8

  x += y    // x = x + y
  fmt.Println(x)  // Output: 50
  
  y -= x    // y = y - x
  fmt.Println(y)  // Output: -42
  
  x /= y    // x = x / y
  fmt.Println(x)  // Output: -1
  
  x *= y    // x = x * y
  fmt.Println(x)  // Output: 42
}
```

## Constants

Constants are variables whose values cannot be changed after they are declared.

In Go, the `const` keyword is used to declare constants.

```Go
package main

import "fmt"

func main() {
  const Pi = 3.14156
  const name = "Emdadul"
  fmt.Println(Pi)    // Output: 3.14156
  fmt.Println(name)  // Output: Emdadul
}
```

## If-Else Statements

If-else statements are used to make decisions in your code based on conditions.

Go supports various comparison operators:
- `==` - Equal to
- `!=` - Not equal to
- `<` - Less than
- `>` - Greater than
- `<=` - Less than or equal to
- `>=` - Greater than or equal to

```Go
package main

import "fmt"

func main() {
    var x int = 500
    
    if x == 500 {
        fmt.Println("Equal")
    }
    
    if x > 200 {
        fmt.Println("x is greater than 200")
    }
    
    if x < 600 {
        fmt.Println("x is less than 600")
    }
}
```

## For Loops

For loops are used to repeat a block of code multiple times.

Go has only one type of loop - the `for` loop, which can be used in different ways.

```Go
package main

import "fmt"

func main() {
   for i := 0; i <= 100; i++ {
       fmt.Println("Go is fun")
   }
}
```

## Functions

Functions are reusable blocks of code that perform a specific task.

In Go, functions are declared using the `func` keyword.

```Go
package main

import "fmt"

func add(a int, b int) {
    x := a + b
    fmt.Println(x)
}

func main() {
  var a, b int
  fmt.Scanln(&a)
  fmt.Scanln(&b)
  
  add(a, b)
  add(5, 78)
}
```

## Functions with Return Values

Functions can return values using the `return` statement.

You need to specify the return type in the function declaration.

```Go
package main

import "fmt"

func sum(a, b, c int) int {
    return a + b + c
}

func main() {
  result := sum(42, 8, 67)
  fmt.Println(result)  // Output: 117
}
```

## Multiple Return Values

Go functions can return multiple values, which is a unique feature of the language.

```Go
package main

import "fmt"

func swap(a, b int) (int, int) {
    return b, a
}

func main() {
  a, b := swap(42, 8)
  fmt.Println(a, b)  // Output: 8 42
  fmt.Println(b)     // Output: 42
}
```

## Structs

Structs are user-defined types that group together variables of different types.

They are similar to classes in other programming languages but without inheritance.

```Go
package main

import "fmt"

type Student struct {
    name string
    age int 
    section string 
    id string
}

func main() {
  X := Student{"Emdadul", 20, "A", "C201041"}
  Y := Student{"Minhaz", 20, "C", "C211091"}
  
  fmt.Println(X.name)     // Output: Emdadul
  fmt.Println(X.age)      // Output: 20
  fmt.Println(X.section)  // Output: A
  fmt.Println(X.id)       // Output: C201041
  
  fmt.Println(Y.name)     // Output: Minhaz
  fmt.Println(Y.age)      // Output: 20
  fmt.Println(Y.section)  // Output: C
  fmt.Println(Y.id)       // Output: C211091
}
```

## Methods

Methods are functions that are associated with a specific type (usually a struct).

Methods allow you to define behavior for your custom types.

```Go
package main

import "fmt"

type Student struct {
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
  X := Student{"Emdadul", "C201041", 20}
  Y := Student{"Minhaz", "C211091", 20}
  
  X.show()  // Calls the show method on X
  Y.show()  // Calls the show method on Y
}
```

## Pointers

Pointers are variables that store the memory address of another variable.

They are useful for passing data by reference and modifying values in functions.

```Go
package main

import "fmt"

func change(val int) {
  val = 8  // This won't change the original value
}

func change_ptr(ptr *int) {
  *ptr = 8  // This will change the original value
}

func main() {
  x := 42

  change(x)
  fmt.Println(x)  // Output: 42 (unchanged)

  change_ptr(&x)
  fmt.Println(x)  // Output: 8 (changed)
}
```

## Slices

Slices are dynamic arrays in Go. They are more flexible than arrays and can grow or shrink.

```Go
package main

import "fmt"

func main() {
  a := [5]int{0, 2, 4, 6, 8}  // Array
  s := a[:4]                  // Slice from array
  fmt.Println(s)              // Output: [0 2 4 6]
}
```

## Range

The `range` keyword is used to iterate over arrays, slices, maps, and strings.

```Go
package main

import "fmt"

func main() {
  a := make([]int, 50)
  a[1] = 2
  a[2] = 3
  a[0] = 78

  for i, v := range a {
      fmt.Println(i+1, v)
  }
}
```

## String Operations

### String Concatenation

Strings can be concatenated using the `+` operator.

```Go
package main

import "fmt"

func main() {
  x := "Emdadul"
  y := "Islam"
  fmt.Println(x + " " + y)  // Output: Emdadul Islam
}
```

### String Range

You can iterate over individual characters in a string using range.

```Go
package main

import "fmt"

func main() {
  x := "Emdadul"
  for _, c := range x {
    fmt.Printf("%c ", c)  // Output: E m d a d u l
  }
}
```

## Defer Statement

The `defer` statement postpones the execution of a function until the surrounding function returns.

Deferred calls are executed in LIFO (Last In, First Out) order.

```Go
package main

import "fmt"

func welcome() {
  fmt.Println("Welcome")
}

func main() {
  defer fmt.Print("hello")  // Executed last
  defer welcome()           // Executed second
  fmt.Println("Hey")        // Executed first
  
  // Output:
  // Hey
  // Welcome
  // hello
}
```

## Make Function

The `make` function is used to create slices, maps, and channels.

For slices, it creates a slice with a specified length and capacity.

```Go
package main

import "fmt"

func main() {
  a := make([]int, 10)  // Creates a slice of length 10
  a = append(a, 8)      // Appends 8 to the slice
  fmt.Println(a)        // Output: [0 0 0 0 0 0 0 0 0 0 8]
}
```

## Append Function

The `append` function is used to add elements to a slice.

```Go
package main

import "fmt"

func main() {
  a := make([]int, 5)
  a = append(a, 8, 4, 6, 7)  // Appends multiple values
  fmt.Println(a)             // Output: [0 0 0 0 0 8 4 6 7]
}
```

## Time Package

The `time` package provides functionality for measuring and displaying time.

```Go
package main

import (
  "fmt"
  "time"
)

func main() {
  fmt.Println(time.Now())  // Shows current time
}
```

## Variable Declaration Short Way

Go supports a short way to declare and initialize multiple variables at once.

```Go
package main

import "fmt"

func main() {
  name, age, salary := "John", 34, 3455.233
  fmt.Println(name)    // Output: John
  fmt.Println(age)     // Output: 34
  fmt.Println(salary)  // Output: 3455.233
}
```
