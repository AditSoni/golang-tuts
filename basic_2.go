package main

import (
	"fmt"
	// "runtime"
)

// runtime tells us about os and cpu arch

func main() {

	// fmt.Println(runtime.GOARCH,runtime.GOOS)

	// s := "Hi kids do you like violence?"

	// fmt.Println(s)
	// fmt.Printf("%T\n",s)

	// bs := []byte(s)

	// fmt.Println(bs)
	// fmt.Printf("%T\n",bs)
	test1()
}

// iota can be used in const declaration only
// each use of iota increments the value by 1 , starts at 0
// there is also implicit repetition that can be used in the language go
// consider the below function where the operations
// 1 << iota and 1 << iota-1 are being done for the next following variables as well

func test1() {
	const (
		bit0, mask0 = 1 << iota, 1 << iota - 1 // bit0 == 1, mask0 == 0  (iota == 0)
		bit1, mask1                          // bit1 == 2, mask1 == 1  (iota == 1)
		_, _                                 //                        (iota == 2, unused)
		bit3, mask3                          // bit3 == 8, mask3 == 7  (iota == 3)
	)

	fmt.Println(bit0, bit1, bit3)
	fmt.Println(mask0, mask1, mask3)
}
