package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstname string
	lastname  string
	contact   contactInfo
}

func main() {
	// // one way, but not recommended
	// alex1 := person{"Alex", "Anderson"}

	// // option 2
	// alex2 := person{firstname: "Alex", lastname: "Anderson2"}

	// // option 3
	// var alex3 person

	// composite struct
	alex4 := person{
		firstname: "Vikas",
		lastname:  "Rajoria",
		contact: contactInfo{
			email:   "v@v.com",
			zipCode: 201005,
		},
	}
	alex4.updateFirstName("Vikas rockstar")

	s := []string{"a", "b", "c"}
	fmt.Println(&s)
	s[0] = "z"
	fmt.Println(&s)
	fmt.Println(s)
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateFirstName(newFirstName string) {
	// jimAdd := &p
	// fmt.Println(jimAdd)
	// (*jimAdd).firstname = newFirstName
	// fmt.Println(*jimAdd)
	// fmt.Println(p)
}
