package main

import "fmt"

func main() {
	// test1()
	// test2()
	// test3()
	// test4()
	// test5()
	// test6()
	// test7()
	test8_9()
}

func test1() {
	arr := [5]int{1, 2, 3, 4, 5}

	for i, v := range arr {
		fmt.Printf("%v %v %T\n", i, v, v)
	}

	fmt.Printf("%T\n", arr)
}

func test2() {
	slicee := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i, v := range slicee {
		fmt.Printf("%v %v %T\n", i, v, v)
	}

	fmt.Printf("%T\n", slicee)
}

func test3() {
	slicee := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(slicee[:5])
	fmt.Println(slicee[5:])
	fmt.Println(slicee[2:7])
	fmt.Println(slicee[1:6])
}
func test4() {
	slicee := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	slicee = append(slicee, 52)
	fmt.Println(slicee)

	slicee = append(slicee, 53, 54, 55)
	fmt.Println(slicee)

	y := []int{56, 57, 58, 59, 60}
	slicee = append(slicee, y...)
	fmt.Println(slicee)

}

func test5() {
	slicee := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	slicee = append(slicee[:3], slicee[6:]...)

	fmt.Println(slicee)

}

func test6() {
	arr := make([]string, 0, 0)
	fmt.Println(len(arr))
	fmt.Println(cap(arr))

	arr = []string{"adit", "adit", "adit", "adit", "adit", "adit", "adit", "adit", "adit", "adit", "adit", "ada", "ada"}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("item at index %v is %v\n", i, arr[i])
	}
	arr = append(arr, "adad")
	fmt.Println(len(arr))
	fmt.Println(cap(arr))
}
func test7() {

	x := [][]string{
		{"james", "bond", "shaken ,not stirred"},
		{"Miss", "moneypenny", "hello,james"},
	}

	for i, sxs := range x {
		fmt.Println(i)
		for j, xxs := range sxs {
			fmt.Println(j, xxs)
		}
	}
}
func test8_9() {
	x := map[string][]string{
		"bond_james":      {"Marini", "Women", "life"},
		"moneypenny_miss": {"hi", "i'm", "dumb"},
		"no_dr":           {"evil", "squirrel", "monkey"},
	}

	for k, v := range x {
		fmt.Println("Key ", k)
		for i, value := range v {
			fmt.Printf("Item at index position %v is %v\n", i, value)
		}
	}

	x["new"] = []string{"Hi", "why", "sigh"}

	for k, v := range x {
		fmt.Println(k, v)
	}

	fmt.Println("DELETEE")

	delete(x, "new")
	for k, v := range x {
		fmt.Println(k, v)
	}
}
