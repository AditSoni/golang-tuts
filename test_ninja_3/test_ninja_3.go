package main

import "fmt"

func main() {
	// test1()
	// test2()
	// test3()
	// test4()
	// test5()
	// test6_7()
	// test8()
	test9()
}

func test1() {

	for x := 1; x <= 10_000; x++ {
		fmt.Println(x)
	}
}

func test2() {

	for i := 65; i < 91; i++ {
		fmt.Println("Loop number ", i-64)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\t%#U\n", i)
		}
	}
}

func test3() {
	x := 1998
	for x <= 2023 {
		fmt.Println(x)
		x++
	}
}

func test4() {
	x := 1998
	for {
		if x > 2023 {
			break
		}
		fmt.Println(x)
		x++
	}
}

func test5() {
	for x := 10; x <= 100; x++ {
		fmt.Printf("%d %% %d = %d\n", x, 4, x%4)
	}
}

func test6_7() {

	if x := 47; x%2 == 0 {
		fmt.Println("222")
	} else if x%3 == 0 {
		fmt.Println("3333")
	} else {
		fmt.Println("nah")
	}
}

func test8() {

	switch {
	case 1 < 2:
		fmt.Println("case 1")
		fallthrough
	default:
		fmt.Println("default")
	case 5 > 6:
		fmt.Println("case 3")
	}
}

func test9() {
	favSport := "abc"
	switch favSport {
	case "bcd":
		fmt.Println("caase bcd")

	case "abc":
		fmt.Println("case abc")

	}
}
