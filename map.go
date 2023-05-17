package main

import "fmt"

func main() {

	mapss()

}

func mapss() {

	// maps are key value pairs (dictionary of python)

	// map is the keyword used , followed by the type of key would be have in []
	// last would be the type of values stored
	x := map[string]int{
		"dev": 10_000,
		"job": 20,
	}

	x["adit"] = 99
	fmt.Println(x)
	// what if a certain value doesn't exist in map ?
	fmt.Println(x["soni"])
	// it prints out the zero value of that type of value

	// we use comma ok ( , ok ) identifier to check if we had the key or not

	v, ok := x["soni"]

	fmt.Println(v, ok)
	// here the second variable ok is bool , to check whether we had this key or not
	// first is init , second is condition
	if v, ok := x["adit"]; ok {
		fmt.Println("Yes this key is  present", v)
	}

	abc, ok := x["adit1"]

	fmt.Println(abc, ok)

	for key, value := range x {
		fmt.Println(key, value)
	}
}
