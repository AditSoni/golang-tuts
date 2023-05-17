package main

import "fmt"

func main() {

	// switch_case()
	// for_loop()
	if_else()
}

func for_loop() {

	// for init ; condition ; post{}
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	// for cond {}
	a := 1
	for a < 3 {
		fmt.Println(a)
		a++
	}

	for letter := 33; letter <= 133; letter++ {
		fmt.Printf("number  %d represents %#U\n", letter, letter)
	}
}

func switch_case() {

	switch {
	case 1 > 2, 1 == 1:
		fmt.Println("case 1 ")
	case 2 >= 2:
		fmt.Println("case 2")
		fallthrough // the next case would run irrespective of it being true.
	case 3 > 5:
		fmt.Println("case 3 ")
	case 4 == 4:
		fmt.Println("case 4 ")
	default: // default if not case is true.
		fmt.Println("default")
	}

	// switch can be on a value as well
	// It would do comparisons with other values
	// here it would check equal value with n
	// can define multiple values in cases as well.
	// multiple conditionals can be placed and would be evalued
	// as if their was an or in between them.
	n := "bond"
	switch n {
	case "a", "b", "bond":
		fmt.Println("Nah 4 ")
	case "v":
		fmt.Println("nah 5 ")
	case "no":
		fmt.Println("yesss")
	}
}

func if_else() {

	// intialization statement
	// scope of x is within if only
	if x := 42; x > 50 {
		fmt.Println(x)
	} else if 1 == 1 {
		fmt.Println(1)
	}
}
