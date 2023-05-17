package main

import "fmt"

// can be declared outside of function scope is the complete package
var z = 99

// can also do type casting to make it store a certain datatype
var z1 string = "123"

// string can be declare between `` also , they are called back ticks

var a = `hi
selene`

// type declaration hotdog is a custom type with underlying type as int
// can use type conversion , there is not concept of casting in go
// var a int
// var b hotdog
// a = int(b) 
// this way we can do type conversions

type hotdog int


func main() {
	// Declare variable only if you want to use it else go will throw error

	// declare and assign value at same time
	// can only be used in function body not outside in global scope ,
	// scope within function body
	x := 44
	fmt.Print("hello ")
	fmt.Println(x)

	// can use to declare variable within the function body but also in global scope
	var y = "Hi dude!"

	fmt.Println(y)
	fmt.Println(z)
	fmt.Println("Different Print Functions : ")
	fmt.Println(z1, x, y)                           //appends new line by default also space after each arg
	fmt.Print("x = ", x, " y = ", y, " ", z1, "\n") // normal print fucntion
	// uses string formatting.
	fmt.Printf("x = %d\a", x)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	type this int
	var axs this

	axs = 11
	fmt.Println(axs)
}
