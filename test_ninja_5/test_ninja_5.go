package main

import "fmt"

func main() {
	// test1()
	// test2()
	// test3()
	test4()
}

type person struct {
	first_name        string
	last_name         string
	ice_cream_flavour []string
}

func test1() {
	p1 := person{
		first_name:        "Adit",
		last_name:         "Soni",
		ice_cream_flavour: []string{"Butter Scotch", "Chocolate", "Cookie and cream"},
	}

	p2 := person{
		first_name:        "Jai",
		last_name:         "Singh",
		ice_cream_flavour: []string{"Vanilla", "Chocolate", "Cookie and cream"},
	}
	fmt.Println("Person 1")
	fmt.Println("\t", p1.first_name, p1.last_name)
	fmt.Println("\tIce cream flavours:")
	for i, v := range p1.ice_cream_flavour {
		fmt.Printf("\t\t\t%v %v\n", i, v)
	}

	fmt.Println("Person 2")
	fmt.Println("\t", p2.first_name, p2.last_name)
	fmt.Println("\tIce cream flavours:")
	for i, v := range p2.ice_cream_flavour {
		fmt.Printf("\t\t\t%v %v\n", i+1, v)
	}
}

func test2() {
	p1 := person{
		first_name:        "Adit",
		last_name:         "Soni",
		ice_cream_flavour: []string{"Butter Scotch", "Chocolate", "Cookie and cream"},
	}

	p2 := person{
		first_name:        "Jai",
		last_name:         "Singh",
		ice_cream_flavour: []string{"Vanilla", "Chocolate", "Cookie and cream"},
	}

	map1 := map[string]person{
		p1.last_name: p1,
		p2.last_name: p2,
	}

	for i, v := range map1 {
		fmt.Println(i, v)
	}

}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func test3() {
	truck1 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "Black",
		},
		fourWheel: true,
	}

	sedan1 := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "Beige",
		},
		luxury: true,
	}

	fmt.Println(truck1)
	fmt.Println(truck1.doors, truck1.color, truck1.fourWheel)

	fmt.Println(sedan1)
	fmt.Println(sedan1.doors, sedan1.color, sedan1.luxury)

}

func test4() {
	s := struct {
		map1   map[string]int
		slice1 []int
	}{
		map1: map[string]int{
			"a": 65,
			"b": 66,
		},
		slice1: []int{1, 2, 3, 4, 5},
	}

	fmt.Println(s)
}
