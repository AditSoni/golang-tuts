package main

import "fmt"

type adit int

var x adit

var y int

func main(){
	fmt.Printf("x = %v type(x) = %T\n",x,x)
	x = 42
	fmt.Println(x)

	y = int(x)

	fmt.Println(y)
	fmt.Printf("%T\n",y)
}