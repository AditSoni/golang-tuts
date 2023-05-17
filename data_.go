package main

import "fmt"

func main() {
	// arrays()
	// slices()
	//make_slice()
	multi_dimensional_slice()
}

func arrays() {
	// declaration of an array
	// 10 is the size and int is the type
	// array are of fixed sizes and we cant add more elements to it.
	// x:= [4]int{1,2,3,4} composite literal
	var x [10]int
	fmt.Println(x)
	x[2] = 11
	fmt.Println(x)

}
// difference between array and slice is we dont specify size for slices
func slices() {
	// slice allows us to group together values of same type
	// composite literal
	x := []int{1, 3, 4, 5}
	fmt.Println(x)

	// looping over a slice
	fmt.Println()
	for i, v := range x {
		fmt.Println(i, v)
	}
	// len is the length of the total elements
	// cap is the capacity ,capacity can be 10 but we used only 5 elements to fill the slice
	// in case of array cap and len are same
	// a slice is a reference type to array
	// array is a value-type
	// slice can increase its capacity as required.
	fmt.Printf("The length is %d  and Capacity is %d before adding elemets\n", len(x), cap(x))
	// can append any number of values of same type
	x = append(x, 40, 51, 63)

	fmt.Printf("The length is %d  and Capacity is %d After adding elemets\n", len(x), cap(x))

	// looping over slice
	for i := 0; i < len(x); i++ {
		fmt.Printf("x[%d] = %d\n", i, x[i])
	}
	// slicing a slice

	fmt.Println("Slicing a slice")
	fmt.Println(x[1:6])

	// append(slice, ...T) can take any number of arguements

	y := []int{444, 555}
	// here y... unfolds or unpacks all of the values of y slice which are used as
	// arguements to append the elements to the slice x
	x = append(x, y...)

	fmt.Printf("The value of x after appending y is %d\n", x)

	// deleting from a slice

	x = append(x[:3], x[5:]...)

	fmt.Println("value of x after deleting elements", x)

}

func make_slice() {
	// would be making slices using make function

	var x []string = make([]string, 10, 12)
	// x := make([]int,0,10)
	// make(type,len,capacity)
	fmt.Println(x)
	x[0], x[1], x[2] = "Hi", "I am", "John"
	fmt.Println(x)

}

func multi_dimensional_slice() {
	people := []string{"Adit", "Soni"}
	food := []string{"Paneer", "Legumes"}

	fmt.Println(people)
	fmt.Println(food)

	fmt.Println("Making a 2 d slice")

	joined := [][]string{people, food}
	fmt.Println(joined)

	fmt.Println(joined[1][0])

	new_2d_slice := [][]string{{"ADIT","Panneer"},{"SONI","Legumes"}}

	fmt.Println(new_2d_slice)

}
