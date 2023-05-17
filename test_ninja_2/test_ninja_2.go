package main

import "fmt"

func main() {
	// test1()
	// test3()
	// test4()
	// test5()
	test6()
}

func test1() {

	x := 22
	fmt.Printf("decimal  = %d\tbinary = %#b\thexa = %#x\n", x, x, x)
}

// test 2 was just conditional operators
func test3() {
	const (
		a        = 5
		b string = "hi"
	)

	fmt.Println(a, b)
}

func test4() {

	x_4 := 42

	fmt.Printf("%d %b %#x\n", x_4, x_4, x_4)
	y_4 := x_4 << 1
	fmt.Printf("%d %b %#x\n", y_4, y_4, y_4)

}

func test5() {
	var x_5 string = `abcd
	"hust" he said`

	fmt.Println(x_5)

}

func test6() {
	const (
		year_1 = 2023 + iota
		year_2 = 2023 + iota
		year_3 = 2023 + iota
		year_4 = 2023 + iota
	)

	fmt.Println(year_1, year_2, year_3, year_4)
}
