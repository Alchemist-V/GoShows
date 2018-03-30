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

	fmt.Println(alex4)
	fmt.Printf("%+v", alex4)

}
