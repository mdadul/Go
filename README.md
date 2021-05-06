1. [Go](#go)
2. [Hello world](#hello-world)
3. [Package](#package)
4. [Import](#import)
5. [Variable](#variables)

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

