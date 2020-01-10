package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

//* embedding a struct inside of another struct can be done with the shorthand
//* instead 'contact contactInfo', you can use 'contactInfo'
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	// * shorthand version of updating a struct without explicitly declaring the pointer value
	// * ie. pointerOfJim := &jim
	// * pointerOfJim.updateName("Nelson")
	jim.updateName("Nelson")
	jim.print()
}

//* '*person' is a pointer type that points to the person 'jim' in memory
func (pointerToPerson *person) updateName(newName string) {
	(*pointerToPerson).firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
