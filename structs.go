package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

// created a type person with struct as underlying type.
// structs can be used to group together data of different types.
// here we used string and int types.

type human struct {
	person
	doa bool
}

// embedded structs
// here sort of we can se we have inheritance where we
// are able to use types inside person for human struct too
func main() {
	p1 := person{
		first: "Adit",
		last:  "Soni",
		age:   25,
	}
	p2 := person{
		first: "Damn",
		last:  "singh",
		age:   24,
	}

	ha1 := human{
		person: person{
			first: "Sand",
			last:  "Man",
			age:   32,
		},
		doa: true,
	}
	fmt.Println(p1, p2)

	fmt.Println(p1.first, p1.last, p1.age)

	fmt.Printf("%T\n", p1)
	fmt.Println(ha1.first, ha1.last, ha1.age, ha1.doa)

	anonymous_struct()

}

func anonymous_struct() {
	p2 := struct {
		subject string
		grade   int
	}{
		subject: "ENglish",
		grade:   100,
	}

	fmt.Println(p2)
	fmt.Printf("%T\n",p2)
}
